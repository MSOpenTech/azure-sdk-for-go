// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package emulation

import (
	"context"
	"fmt"
	"log"
	"math"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/amqpwrap"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/auth"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/go-amqp"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/mock"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus/internal/sbauth"
	"github.com/golang/mock/gomock"
)

type MockData struct {
	nextID int64
	Ctrl   *gomock.Controller

	Events *Events

	cbsRouterOnce sync.Once
	options       *MockDataOptions

	queuesMu sync.Mutex
	queues   map[string]*Queue
}

type MockDataOptions struct {
	NewConnectionFn func(orig *mock.MockAMQPClient, ctx context.Context) error

	// PreReceiverMock is called with the mock Receiver instance before we instrument it with
	// our defaults.
	PreReceiverMock func(mr *MockReceiver, ctx context.Context) error

	// PreSenderMock is called with the mock Sender instance before we instrument it with
	// our defaults.
	PreSenderMock func(ms *MockSender, ctx context.Context) error

	// PreSessionMock is called with the mock Session instance before we instrument it with
	// our defaults.
	PreSessionMock func(msess *MockSession, ctx context.Context, opts *amqp.SessionOptions) error
}

func NewMockData(t *testing.T, options *MockDataOptions) *MockData {
	if options == nil {
		options = &MockDataOptions{}
	}

	if options.PreReceiverMock == nil {
		options.PreReceiverMock = func(orig *MockReceiver, ctx context.Context) error {
			return nil
		}
	}

	if options.PreSenderMock == nil {
		options.PreSenderMock = func(orig *MockSender, ctx context.Context) error {
			return nil
		}
	}

	if options.PreSessionMock == nil {
		options.PreSessionMock = func(orig *MockSession, ctx context.Context, opts *amqp.SessionOptions) error {
			return nil
		}
	}

	if options.NewConnectionFn == nil {
		options.NewConnectionFn = func(orig *mock.MockAMQPClient, ctx context.Context) error {
			return nil
		}
	}

	return &MockData{
		Ctrl:    gomock.NewController(t),
		queues:  map[string]*Queue{},
		Events:  &Events{},
		options: options,
	}
}

type MockConnection struct {
	Status *Status
	*mock.MockAMQPClient
}

type MockSession struct {
	Status *Status
	*mock.MockAMQPSession

	ID   string
	Conn *MockConnection
}

func (sess *MockSession) Done() <-chan error {
	errCh := make(chan error)

	go func() {
		select {
		case <-sess.Conn.Status.Done():
			errCh <- sess.Conn.Status.Err()
		case <-sess.Status.Done():
			errCh <- sess.Status.Err()
		}
	}()

	return errCh
}

func (md *MockData) NewConnection(ctx context.Context) (amqpwrap.AMQPClient, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	conn := &MockConnection{
		MockAMQPClient: mock.NewMockAMQPClient(md.Ctrl),
		Status:         NewStatus(nil),
	}

	connID := md.nextUniqueName("c")

	conn.EXPECT().Name().Return(connID).AnyTimes()
	conn.EXPECT().NewSession(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, opts *amqp.SessionOptions) (amqpwrap.AMQPSession, error) {
		return md.newSession(ctx, opts, conn)
	}).AnyTimes()

	conn.EXPECT().Close().DoAndReturn(func() error {
		md.Events.CloseConnection(conn.Name())
		conn.Status.CloseWithError(&amqp.ConnectionError{})
		return nil
	}).AnyTimes()

	md.Events.OpenConnection(conn.Name())
	return conn, nil
}

func (md *MockData) newSession(ctx context.Context, opts *amqp.SessionOptions, conn *MockConnection) (amqpwrap.AMQPSession, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-conn.Status.Done():
		return nil, conn.Status.Err()
	default:
	}

	sess := &MockSession{
		Status:          NewStatus(conn.Status),
		ID:              md.nextUniqueName("sess"),
		Conn:            conn,
		MockAMQPSession: mock.NewMockAMQPSession(md.Ctrl),
	}

	if err := md.options.PreSessionMock(sess, ctx, opts); err != nil {
		return nil, err
	}

	sess.EXPECT().NewReceiver(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, source string, opts *amqp.ReceiverOptions) (amqpwrap.AMQPReceiverCloser, error) {
		return md.NewReceiver(ctx, source, opts, sess)
	}).AnyTimes()

	sess.EXPECT().NewSender(gomock.Any(), gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, target string, opts *amqp.SenderOptions) (amqpwrap.AMQPSenderCloser, error) {
		return md.NewSender(ctx, target, opts, sess)
	}).AnyTimes()

	sess.EXPECT().Close(gomock.Any()).DoAndReturn(func(ctx context.Context) error {
		select {
		case <-conn.Status.Done():
			return conn.Status.Err()
		default:
			sess.Status.CloseWithError(amqp.ErrSessionClosed)
			return nil
		}
	}).AnyTimes()

	return sess, nil
}

func (md *MockData) upsertQueue(name string) *Queue {
	md.queuesMu.Lock()
	defer md.queuesMu.Unlock()

	q := md.queues[name]

	if q == nil {
		q = NewQueue(name, md.Events)
		md.queues[name] = q
	}

	return q
}

func (md *MockData) getQueue(name string) *Queue {
	md.queuesMu.Lock()
	defer md.queuesMu.Unlock()

	return md.queues[name]
}

func (md *MockData) AllQueues() map[string]*Queue {
	md.queuesMu.Lock()
	defer md.queuesMu.Unlock()

	m := map[string]*Queue{}

	for k, v := range md.queues {
		m[k] = v
	}

	return m
}

func (md *MockData) nextUniqueName(prefix string) string {
	nextID := atomic.AddInt64(&md.nextID, 1)
	return fmt.Sprintf("%s:%03X", prefix, nextID)
}

func (md *MockData) NewTokenProvider() auth.TokenProvider {
	tc := mock.NewMockTokenCredential(md.Ctrl)

	var tokenCounter int64

	tc.EXPECT().GetToken(gomock.Any(), gomock.Any()).DoAndReturn(func(ctx context.Context, options policy.TokenRequestOptions) (azcore.AccessToken, error) {
		tc := atomic.AddInt64(&tokenCounter, 1)

		return azcore.AccessToken{
			Token:     fmt.Sprintf("Token-%d", tc),
			ExpiresOn: time.Now().Add(10 * time.Minute),
		}, nil
	}).AnyTimes()

	return sbauth.NewTokenProvider(tc)
}

func (md *MockData) cbsRouter(ctx context.Context, in *Queue, getQueue func(name string) *Queue) {
	_ = in.IssueCredit(math.MaxUint32, LinkEvent{
		ConnID: "none",
		Entity: in.name,
		Name:   "none",
		Role:   LinkRoleReceiver,
	}, nil)

	for {
		msg, err := in.Receive(ctx, LinkEvent{
			ConnID: "none",
			Entity: in.name,
			Name:   "none",
			Role:   LinkRoleReceiver,
		}, nil)

		if err != nil {
			log.Printf("CBS router closed: %s", err.Error())
			break
		}

		// route response to the right spot
		replyTo := *msg.Properties.ReplyTo

		out := getQueue(replyTo)

		// assume auth is valid.
		err = out.Send(ctx, &amqp.Message{
			Properties: &amqp.MessageProperties{
				CorrelationID: msg.Properties.MessageID,
			},
			ApplicationProperties: map[string]any{
				"statusCode": int32(200),
			},
		}, LinkEvent{
			ConnID: "none",
			Entity: out.name,
			Name:   "none",
			Role:   LinkRoleSender,
		}, nil)

		if err != nil {
			log.Printf("CBS router closed: %s", err.Error())
			break
		}
	}
}
