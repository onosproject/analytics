package kafkaClient

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func StartTopicReader(ctx context.Context, messageChannel chan string, errorChannel chan error, brokerURLs []string, inbound string, groupID string) {

	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokerURLs,
		Topic:   inbound,
		GroupID: groupID,
	})
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			errorChannel <- err
			continue
		}
		messageChannel <- string(msg.Value)
	}
}
