// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package azservicebus

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal"
	"github.com/Azure/go-amqp"
	"github.com/devigned/tab"
)

// ReceiveMode represents the lock style to use for a receiver - either
// `PeekLock` or `ReceiveAndDelete`
type ReceiveMode = internal.ReceiveMode

const (
	// PeekLock will lock messages as they are received and can be settled
	// using the Receiver or Processor's (Complete|Abandon|DeadLetter|Defer)Message
	// functions.
	PeekLock ReceiveMode = internal.PeekLock
	// ReceiveAndDelete will delete messages as they are received.
	ReceiveAndDelete ReceiveMode = internal.ReceiveAndDelete
)

// SubQueue allows you to target a subqueue of a queue or subscription.
// Ex: the dead letter queue (SubQueueDeadLetter).
type SubQueue int

const (
	// SubQueueDeadLetter targets the dead letter queue for a queue or subscription.
	SubQueueDeadLetter SubQueue = 1
	// SubQueueTransfer targets the transfer dead letter queue for a queue or subscription.
	SubQueueTransfer SubQueue = 2
)

// Receiver receives messages using pull based functions (ReceiveMessages).
// For push-based receiving via callbacks look at the `Processor` type.
type Receiver struct {
	settler settler

	config receiverConfig

	lastPeekedSequenceNumber int64
	amqpLinks                internal.AMQPLinks

	mu        sync.Mutex
	receiving bool
}

type receiverConfig struct {
	ReceiveMode    ReceiveMode
	FullEntityPath string
	Entity         entity

	baseRetrier    internal.Retrier
	cleanupOnClose func()
}

const defaultLinkRxBuffer = 2048

// ReceiverOption represents an option for a receiver.
// Some examples:
// - `ReceiverWithReceiveMode` to configure the receive mode,
// - `ReceiverWithSubQueue` to connect to a subqueue, like a dead letter queue
type ReceiverOption func(receiver *Receiver) error

// ReceiverWithSubQueue allows you to open the sub queue (ie: dead letter queues, transfer dead letter queues)
// for a queue or subscription.
func ReceiverWithSubQueue(subQueue SubQueue) ReceiverOption {
	return func(r *Receiver) error {
		return r.config.Entity.SetSubQueue(subQueue)
	}
}

// ReceiverWithReceiveMode controls the receive mode for the receiver.
func ReceiverWithReceiveMode(receiveMode ReceiveMode) ReceiverOption {
	return func(receiver *Receiver) error {
		if receiveMode != PeekLock && receiveMode != ReceiveAndDelete {
			return fmt.Errorf("invalid receive mode specified %d", receiveMode)
		}

		receiver.config.ReceiveMode = receiveMode
		return nil
	}
}

func newReceiver(ns internal.NamespaceWithNewAMQPLinks, cleanupOnClose func(), options ...ReceiverOption) (*Receiver, error) {
	receiver := &Receiver{
		config: receiverConfig{
			ReceiveMode: PeekLock,
			// TODO: make this configurable
			baseRetrier: internal.NewBackoffRetrier(internal.BackoffRetrierParams{
				Factor:     1.5,
				Jitter:     true,
				Min:        time.Second,
				Max:        time.Minute,
				MaxRetries: 10,
			}),
			cleanupOnClose: cleanupOnClose,
		},
		lastPeekedSequenceNumber: 0,
	}

	for _, opt := range options {
		if err := opt(receiver); err != nil {
			return nil, err
		}
	}

	entityPath, err := receiver.config.Entity.String()

	if err != nil {
		return nil, err
	}

	receiver.amqpLinks = ns.NewAMQPLinks(entityPath, func(ctx context.Context, session internal.AMQPSession) (internal.AMQPSenderCloser, internal.AMQPReceiverCloser, error) {
		linkOptions := createLinkOptions(receiver.config.ReceiveMode, entityPath)
		return createReceiverLink(ctx, session, linkOptions)
	})

	// 'nil' settler handles returning an error message for receiveAndDelete links.
	if receiver.config.ReceiveMode == PeekLock {
		receiver.settler = newMessageSettler(receiver.amqpLinks, receiver.config.baseRetrier)
	}

	return receiver, nil
}

// receiveOptions are options for the ReceiveMessages function.
type receiveOptions struct {
	maxWaitTime                  time.Duration
	maxWaitTimeAfterFirstMessage time.Duration
}

// ReceiveOption represents an option for a `ReceiveMessages`.
// For example, `ReceiveWithMaxWaitTime` will let you configure the
// maxmimum amount of time to wait for messages to arrive.
type ReceiveOption func(options *receiveOptions) error

// ReceiveWithMaxWaitTime configures how long to wait for the first
// message in a set of messages to arrive.
// Default: 60 seconds
func ReceiveWithMaxWaitTime(max time.Duration) ReceiveOption {
	return func(options *receiveOptions) error {
		options.maxWaitTime = max
		return nil
	}
}

// ReceiveWithMaxTimeAfterFirstMessage confiures how long, after the first
// message arrives, to wait before returning.
// Default: 1 second
func ReceiveWithMaxTimeAfterFirstMessage(max time.Duration) ReceiveOption {
	return func(options *receiveOptions) error {
		options.maxWaitTimeAfterFirstMessage = max
		return nil
	}
}

// ReceiveMessages receives a fixed number of messages, up to numMessages.
// There are two timeouts involved in receiving messages:
// 1. An explicit timeout set with `ReceiveWithMaxWaitTime` (default: 60 seconds)
// 2. An implicit timeout (default: 1 second) that starts after the first
//    message has been received. This time can be adjusted with `ReceiveWithMaxTimeAfterFirstMessage`
func (r *Receiver) ReceiveMessages(ctx context.Context, maxMessages int, options ...ReceiveOption) ([]*ReceivedMessage, error) {
	r.mu.Lock()
	isReceiving := r.receiving

	if !isReceiving {
		r.receiving = true

		defer func() {
			r.mu.Lock()
			r.receiving = false
			r.mu.Unlock()
		}()
	}
	r.mu.Unlock()

	if isReceiving {
		return nil, errors.New("receiver is already receiving messages. ReceiveMessages() cannot be called concurrently.")
	}

	return r.receiveMessagesImpl(ctx, maxMessages, options...)
}

func (r *Receiver) receiveMessagesImpl(ctx context.Context, maxMessages int, options ...ReceiveOption) ([]*ReceivedMessage, error) {
	// There are three phases for this function:
	// Phase 1. <receive, respecting user cancellation>
	// Phase 2. <check error and exit if fatal>
	//    NOTE: We don't exit here so we don't end up buffering messages internally that the
	//    user isn't actually waiting for anymore. So we make sure that #3 runs if the
	//    link is still valid.
	// Phase 3. <drain the link and leave it in a good state>
	ropts := &receiveOptions{
		maxWaitTime:                  time.Minute,
		maxWaitTimeAfterFirstMessage: time.Second,
	}

	for _, opt := range options {
		if err := opt(ropts); err != nil {
			return nil, err
		}
	}

	_, receiver, _, linksRevision, err := r.amqpLinks.Get(ctx)

	if err != nil {
		if err := r.amqpLinks.RecoverIfNeeded(ctx, linksRevision, err); err != nil {
			return nil, err
		}

		return nil, err
	}

	if err := receiver.IssueCredit(uint32(maxMessages)); err != nil {
		_ = r.amqpLinks.RecoverIfNeeded(ctx, linksRevision, err)
		return nil, err
	}

	messages, err := r.getMessages(ctx, receiver, maxMessages, ropts)

	if err != nil {
		return nil, err
	}

	if len(messages) == maxMessages {
		// no drain needed, all messages arrived.
		return messages, nil
	}

	return r.drainLink(receiver, messages)
}

// drainLink initiates a drainLink on the link. Service Bus will send whatever messages it might have still had and
// set our link credit to 0.
func (r *Receiver) drainLink(receiver internal.AMQPReceiver, messages []*ReceivedMessage) ([]*ReceivedMessage, error) {
	receiveCtx, cancelReceive := context.WithCancel(context.Background())

	// start the drain asynchronously. Note that we ignore the user's context at this point
	// since draining makes sure we don't get messages when nobody is receiving.
	go func() {
		if err := receiver.DrainCredit(context.Background()); err != nil {
			tab.For(receiveCtx).Debug(fmt.Sprintf("Draining of credit failed. link will be closed and will re-open on next receive: %s", err.Error()))

			// if the drain fails we just close the link so it'll re-open at the next receive.
			if err := r.amqpLinks.Close(context.Background(), false); err != nil {
				tab.For(receiveCtx).Debug(fmt.Sprintf("Failed to close links on ReceiveMessages cleanup. Not fatal: %s", err.Error()))
			}
		}
		cancelReceive()
	}()

	// Receive until the drain completes, at which point it'll cancel
	// our context.
	// NOTE: That's a gap here where we need to be able to drain _only_ the internally cached messages
	// in the receiver. Filed as https://github.com/Azure/go-amqp/issues/71
	for {
		am, err := receiver.Receive(receiveCtx)

		if internal.IsCancelError(err) {
			break
		} else if err != nil {
			// something fatal happened, we will just
			_ = r.amqpLinks.Close(context.TODO(), false)

			if len(messages) > 0 {
				return messages, nil
			} else {
				return nil, err
			}
		}

		messages = append(messages, newReceivedMessage(receiveCtx, am))
	}

	return messages, nil
}

// getMessages receives messages until a link failure, timeout or the user
// cancels their context.
func (r *Receiver) getMessages(theirCtx context.Context, receiver internal.AMQPReceiver, maxMessages int, ropts *receiveOptions) ([]*ReceivedMessage, error) {
	ctx, cancel := context.WithTimeout(theirCtx, ropts.maxWaitTime)
	defer cancel()

	var messages []*ReceivedMessage

	for {
		var amqpMessage *amqp.Message
		amqpMessage, err := receiver.Receive(ctx)

		if err != nil {
			if internal.IsCancelError(err) {
				return messages, nil
			}

			// we'll close (instead of recovering) since we are holding onto messages
			// and want to get them back to the user ASAP. (recovery will just happen
			// on the next call to receive)
			if err := r.amqpLinks.Close(context.Background(), false); err != nil {
				tab.For(ctx).Debug(fmt.Sprintf("Failed to close links on ReceiveMessages cleanup. Not fatal: %s", err.Error()))
			}
			return nil, err
		}

		messages = append(messages, newReceivedMessage(ctx, amqpMessage))

		if len(messages) == maxMessages {
			return messages, nil
		}

		if len(messages) == 1 {
			go func() {
				select {
				case <-time.After(ropts.maxWaitTimeAfterFirstMessage):
					cancel()
				case <-ctx.Done():
					break
				}
			}()
		}
	}
}

// ReceiveDeferredMessages receives messages that were deferred using `Receiver.DeferMessage`.
func (r *Receiver) ReceiveDeferredMessages(ctx context.Context, sequenceNumbers []int64) ([]*ReceivedMessage, error) {
	_, _, mgmt, _, err := r.amqpLinks.Get(ctx)

	if err != nil {
		return nil, err
	}

	amqpMessages, err := mgmt.ReceiveDeferred(ctx, r.config.ReceiveMode, sequenceNumbers)

	if err != nil {
		return nil, err
	}

	var receivedMessages []*ReceivedMessage

	for _, amqpMsg := range amqpMessages {
		receivedMsg := newReceivedMessage(ctx, amqpMsg)
		receivedMsg.deferred = true

		receivedMessages = append(receivedMessages, receivedMsg)
	}

	return receivedMessages, nil
}

type peekOptions struct {
	fromSequenceNumber *int64
}

type PeekOption func(p *peekOptions) error

// PeekFromSequenceNumber sets the sequence number to start peeking
// messages.
func PeekFromSequenceNumber(sequenceNumber int64) PeekOption {
	return func(p *peekOptions) error {
		p.fromSequenceNumber = &sequenceNumber
		return nil
	}
}

// PeekMessages will peek messages without locking or deleting messages.
// Messages that are peeked do not have lock tokens, so settlement methods
// like CompleteMessage, AbandonMessage, DeferMessage or DeadLetterMessage
// will not work with them.
func (r *Receiver) PeekMessages(ctx context.Context, maxMessageCount int, options ...PeekOption) ([]*ReceivedMessage, error) {

	peekOptions := &peekOptions{}

	for _, opt := range options {
		if err := opt(peekOptions); err != nil {
			return nil, err
		}
	}

	_, _, mgmt, _, err := r.amqpLinks.Get(ctx)

	if err != nil {
		return nil, err
	}

	var sequenceNumber = r.lastPeekedSequenceNumber + 1

	if peekOptions.fromSequenceNumber != nil {
		sequenceNumber = *peekOptions.fromSequenceNumber
	}

	messages, err := mgmt.PeekMessages(ctx, sequenceNumber, int32(maxMessageCount))

	if err != nil {
		return nil, err
	}

	receivedMessages := make([]*ReceivedMessage, len(messages))

	for i := 0; i < len(messages); i++ {
		receivedMessages[i] = newReceivedMessage(ctx, messages[i])
	}

	if len(receivedMessages) > 0 && peekOptions.fromSequenceNumber == nil {
		// only update this if they're doing the implicit iteration as part of the receiver.
		r.lastPeekedSequenceNumber = *receivedMessages[len(receivedMessages)-1].SequenceNumber
	}

	return receivedMessages, nil
}

// receiveDeferredMessage receives a single message that was deferred using `Receiver.DeferMessage`.
func (r *Receiver) receiveDeferredMessage(ctx context.Context, sequenceNumber int64) (*ReceivedMessage, error) {
	messages, err := r.ReceiveDeferredMessages(ctx, []int64{sequenceNumber})

	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, nil
	}

	return messages[0], nil
}

// receiveMessage receives a single message.
func (r *Receiver) receiveMessage(ctx context.Context, options ...ReceiveOption) (*ReceivedMessage, error) {
	messages, err := r.ReceiveMessages(ctx, 1, options...)

	if err != nil {
		return nil, err
	}

	if len(messages) == 0 {
		return nil, nil
	}

	return messages[0], nil
}

// Close permanently closes the receiver.
func (r *Receiver) Close(ctx context.Context) error {
	r.config.cleanupOnClose()
	return r.amqpLinks.Close(ctx, true)
}

// CompleteMessage completes a message, deleting it from the queue or subscription.
func (r *Receiver) CompleteMessage(ctx context.Context, message *ReceivedMessage) error {
	return r.settler.CompleteMessage(ctx, message)
}

// AbandonMessage will cause a message to be returned to the queue or subscription.
// This will increment its delivery count, and potentially cause it to be dead lettered
// depending on your queue or subscription's configuration.
func (r *Receiver) AbandonMessage(ctx context.Context, message *ReceivedMessage) error {
	return r.settler.AbandonMessage(ctx, message)
}

// DeferMessage will cause a message to be deferred. Deferred messages
// can be received using `Receiver.ReceiveDeferredMessages`.
func (r *Receiver) DeferMessage(ctx context.Context, message *ReceivedMessage) error {
	return r.settler.DeferMessage(ctx, message)
}

// DeadLetterMessage settles a message by moving it to the dead letter queue for a
// queue or subscription. To receive these messages create a receiver with `Client.NewReceiver()`
// using the `ReceiverWithSubQueue()` option.
func (r *Receiver) DeadLetterMessage(ctx context.Context, message *ReceivedMessage, options ...DeadLetterOption) error {
	return r.settler.DeadLetterMessage(ctx, message, options...)
}

type entity struct {
	subqueue     SubQueue
	Queue        string
	Topic        string
	Subscription string
}

func (e *entity) String() (string, error) {
	entityPath := ""

	if e.Queue != "" {
		entityPath = e.Queue
	} else if e.Topic != "" && e.Subscription != "" {
		entityPath = fmt.Sprintf("%s/Subscriptions/%s", e.Topic, e.Subscription)
	} else {
		return "", errors.New("a queue or subscription was not specified")
	}

	if e.subqueue == SubQueueDeadLetter {
		entityPath += "/$DeadLetterQueue"
	} else if e.subqueue == SubQueueTransfer {
		entityPath += "/$Transfer/$DeadLetterQueue"
	}

	return entityPath, nil
}

func (e *entity) SetSubQueue(subQueue SubQueue) error {
	if subQueue == SubQueueDeadLetter || subQueue == SubQueueTransfer {
		e.subqueue = subQueue
	} else {
		return fmt.Errorf("unknown SubQueue %d", subQueue)
	}

	return nil
}

func createReceiverLink(ctx context.Context, session internal.AMQPSession, linkOptions []amqp.LinkOption) (internal.AMQPSenderCloser, internal.AMQPReceiverCloser, error) {
	amqpReceiver, err := session.NewReceiver(linkOptions...)

	if err != nil {
		tab.For(ctx).Error(err)
		return nil, nil, err
	}

	return nil, amqpReceiver, nil
}

func createLinkOptions(mode ReceiveMode, entityPath string) []amqp.LinkOption {
	receiveMode := amqp.ModeSecond

	if mode == ReceiveAndDelete {
		receiveMode = amqp.ModeFirst
	}

	opts := []amqp.LinkOption{
		amqp.LinkSourceAddress(entityPath),
		amqp.LinkReceiverSettle(receiveMode),
		amqp.LinkWithManualCredits(),
		amqp.LinkCredit(defaultLinkRxBuffer),
	}

	if mode == ReceiveAndDelete {
		opts = append(opts, amqp.LinkSenderSettle(amqp.ModeSettled))
	}

	return opts
}

// receiverWithQueue configures a receiver to connect to a queue.
func receiverWithQueue(queue string) ReceiverOption {
	return func(receiver *Receiver) error {
		receiver.config.Entity.Queue = queue
		return nil
	}
}

// receiverWithSubscription configures a receiver to connect to a subscription
// associated with a topic.
func receiverWithSubscription(topic string, subscription string) ReceiverOption {
	return func(receiver *Receiver) error {
		receiver.config.Entity.Topic = topic
		receiver.config.Entity.Subscription = subscription
		return nil
	}
}
