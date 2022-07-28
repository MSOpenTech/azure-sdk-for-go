// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.
package azeventhubs_test

import (
	"context"
	"log"
	"sync"
	"testing"
	"time"

	azlog "github.com/Azure/azure-sdk-for-go/sdk/azcore/log"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/to"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azeventhubs"
	"github.com/stretchr/testify/require"
)

func TestConsumerClient_DefaultAzureCredential(t *testing.T) {
	testParams := getConnectionParams(t)

	dac, err := azidentity.NewDefaultAzureCredential(nil)
	require.NoError(t, err)

	t.Run("EventHubProperties and PartitionProperties", func(t *testing.T) {
		consumerClient, err := azeventhubs.NewConsumerClient(testParams.EventHubNamespace, testParams.EventHubName, "0", azeventhubs.DefaultConsumerGroup, dac, nil)
		require.NoError(t, err)

		defer func() {
			err := consumerClient.Close(context.Background())
			require.NoError(t, err)
		}()

		producerClient, err := azeventhubs.NewProducerClient(testParams.EventHubNamespace, testParams.EventHubName, dac, nil)
		require.NoError(t, err)

		defer func() {
			err := producerClient.Close(context.Background())
			require.NoError(t, err)
		}()

		consumerProps, err := consumerClient.GetEventHubProperties(context.Background(), nil)
		require.NoError(t, err)

		producerProps, err := producerClient.GetEventHubProperties(context.Background(), nil)
		require.NoError(t, err)

		require.Equal(t, consumerProps, producerProps)

		producerPartProps, err := producerClient.GetPartitionProperties(context.Background(), consumerProps.PartitionIDs[0], nil)
		require.NoError(t, err)

		consumerPartProps, err := consumerClient.GetPartitionProperties(context.Background(), consumerProps.PartitionIDs[0], nil)
		require.NoError(t, err)

		require.Equal(t, producerPartProps, consumerPartProps)
	})

	t.Run("send and receive", func(t *testing.T) {
		producerClient, err := azeventhubs.NewProducerClient(testParams.EventHubNamespace, testParams.EventHubName, dac, nil)
		require.NoError(t, err)

		defer func() {
			err := producerClient.Close(context.Background())
			require.NoError(t, err)
		}()

		firstPartition, err := producerClient.GetPartitionProperties(context.Background(), "0", nil)
		require.NoError(t, err)

		consumerClient, err := azeventhubs.NewConsumerClient(testParams.EventHubNamespace, testParams.EventHubName, firstPartition.PartitionID, azeventhubs.DefaultConsumerGroup, dac,
			&azeventhubs.ConsumerClientOptions{
				StartPosition: getStartPosition(firstPartition),
			})
		require.NoError(t, err)

		defer func() {
			err := consumerClient.Close(context.Background())
			require.NoError(t, err)
		}()

		eventDataBatch, err := producerClient.NewEventDataBatch(context.Background(), &azeventhubs.NewEventDataBatchOptions{
			PartitionID: to.Ptr(firstPartition.PartitionID),
		})
		require.NoError(t, err)

		err = eventDataBatch.AddEventData(&azeventhubs.EventData{
			Body: []byte("hello"),
		}, nil)
		require.NoError(t, err)

		err = producerClient.SendEventBatch(context.Background(), eventDataBatch, nil)
		require.NoError(t, err)

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		events, err := consumerClient.ReceiveEvents(ctx, 1, nil)
		require.NoError(t, err)
		require.NotEmpty(t, events)
		require.Equal(t, "hello", string(events[0].Body))

		consumerPart, err := consumerClient.GetPartitionProperties(context.Background(), firstPartition.PartitionID, nil)
		require.NoError(t, err)
		producerPart, err := producerClient.GetPartitionProperties(context.Background(), firstPartition.PartitionID, nil)
		require.NoError(t, err)

		require.Equal(t, firstPartition.LastEnqueuedSequenceNumber+1, consumerPart.LastEnqueuedSequenceNumber)
		require.Equal(t, consumerPart, producerPart)
	})

	t.Run("EventHubProperties and PartitionProperties after send", func(t *testing.T) {
		consumerClient, err := azeventhubs.NewConsumerClient(testParams.EventHubNamespace, testParams.EventHubName, "0", azeventhubs.DefaultConsumerGroup, dac, nil)
		require.NoError(t, err)

		defer func() {
			err := consumerClient.Close(context.Background())
			require.NoError(t, err)
		}()

		producerClient, err := azeventhubs.NewProducerClient(testParams.EventHubNamespace, testParams.EventHubName, dac, nil)
		require.NoError(t, err)

		defer func() {
			err := producerClient.Close(context.Background())
			require.NoError(t, err)
		}()

		consumerProps, err := consumerClient.GetEventHubProperties(context.Background(), nil)
		require.NoError(t, err)

		producerProps, err := producerClient.GetEventHubProperties(context.Background(), nil)
		require.NoError(t, err)

		require.Equal(t, consumerProps, producerProps)

		producerPartProps, err := producerClient.GetPartitionProperties(context.Background(), consumerProps.PartitionIDs[0], nil)
		require.NoError(t, err)

		consumerPartProps, err := consumerClient.GetPartitionProperties(context.Background(), consumerProps.PartitionIDs[0], nil)
		require.NoError(t, err)

		require.Equal(t, producerPartProps, consumerPartProps)
	})

}

func TestConsumerClient_GetHubAndPartitionProperties(t *testing.T) {
	testParams := getConnectionParams(t)

	consumer, err := azeventhubs.NewConsumerClientFromConnectionString(testParams.ConnectionString, testParams.EventHubName, "0", azeventhubs.DefaultConsumerGroup, nil)
	require.NoError(t, err)

	defer func() {
		err := consumer.Close(context.Background())
		require.NoError(t, err)
	}()

	hubProps, err := consumer.GetEventHubProperties(context.Background(), nil)
	require.NoError(t, err)
	require.NotEmpty(t, hubProps.PartitionIDs)

	for _, partitionID := range hubProps.PartitionIDs {
		props, err := consumer.GetPartitionProperties(context.Background(), partitionID, nil)
		require.NoError(t, err)

		require.Equal(t, partitionID, props.PartitionID)
	}
}

func TestConsumerClient_Epochs(t *testing.T) {
	azlog.SetEvents(azeventhubs.EventConsumer)
	azlog.SetListener(func(e azlog.Event, s string) {
		log.Printf("[%s] %s", e, s)
	})

	testParams := getConnectionParams(t)

	partitions := mustSendEventToAllPartitions(t, testParams.ConnectionString, testParams.EventHubName, &azeventhubs.EventData{
		Body: []byte("hello world"),
	})

	const concurrentClients = 5 // max you can have with a single consumer group

	clients := make(chan *azeventhubs.ConsumerClient, concurrentClients)

	// this is fine - you can have multiple parallel consumers so long as nobody specifies an epoch
	for i := 0; i < concurrentClients; i++ {
		go func() {
			client, err := azeventhubs.NewConsumerClientFromConnectionString(testParams.ConnectionString, testParams.EventHubName, partitions[0].PartitionID, "$Default", &azeventhubs.ConsumerClientOptions{
				StartPosition: getStartPosition(partitions[0]),
			})
			require.NoError(t, err)

			defer func() {
				err := client.Close(context.Background())
				require.NoError(t, err)
			}()

			ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
			defer cancel()

			events, err := client.ReceiveEvents(ctx, 2, nil)
			require.NoError(t, err)

			require.Equal(t, 1, len(events))
			clients <- client
		}()
	}

	for i := 0; i < concurrentClients; i++ {
		client := <-clients
		require.NoError(t, client.Close(context.Background()))
	}
}

// mustSendEventToAllPartitions sends the event given in evt to each partition in the
// eventHub, returning the sequence number just before the new message.
//
// This is useful for tests that need to work with a hub that might already have messages, and need
// to start from a particular sequence number to avoid them.
//
// NOTE: the message that's passed in does get altered so don't count on it being unchanged after calling
// this function. Each message gets an additional property (DestPartitionID), set to the parttion ID that
// we sent it to.
func mustSendEventToAllPartitions(t *testing.T, cs string, eventHub string, evt *azeventhubs.EventData) []azeventhubs.PartitionProperties {
	producer, err := azeventhubs.NewProducerClientFromConnectionString(cs, eventHub, nil)
	require.NoError(t, err)

	defer func() {
		err := producer.Close(context.Background())
		require.NoError(t, err)
	}()

	hubProps, err := producer.GetEventHubProperties(context.Background(), nil)
	require.NoError(t, err)

	var partitions []azeventhubs.PartitionProperties

	wg := sync.WaitGroup{}
	wg.Add(len(hubProps.PartitionIDs))

	for _, partitionID := range hubProps.PartitionIDs {
		go func(partitionID string) {
			defer wg.Done()

			partProps, err := producer.GetPartitionProperties(context.Background(), partitionID, nil)
			require.NoError(t, err)
			partitions = append(partitions, partProps)

			// send the message to the partition.
			batch, err := producer.NewEventDataBatch(context.Background(), &azeventhubs.NewEventDataBatchOptions{
				PartitionID: &partitionID,
			})
			require.NoError(t, err)

			if evt.ApplicationProperties == nil {
				evt.ApplicationProperties = map[string]any{}
			}

			evt.ApplicationProperties["DestPartitionID"] = partitionID

			err = batch.AddEventData(evt, nil)
			require.NoError(t, err)

			err = producer.SendEventBatch(context.Background(), batch, nil)
			require.NoError(t, err)
		}(partitionID)
	}

	wg.Wait()

	return partitions
}

func getStartPosition(props azeventhubs.PartitionProperties) azeventhubs.StartPosition {
	if props.IsEmpty {
		return azeventhubs.StartPosition{
			Earliest:  to.Ptr(true),
			Inclusive: true,
		}
	}

	return azeventhubs.StartPosition{
		SequenceNumber: to.Ptr(props.LastEnqueuedSequenceNumber),
		Inclusive:      false,
	}
}
