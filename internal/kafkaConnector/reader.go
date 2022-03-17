/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package kafkaConnector

import (
	"context"
	"log"

	"github.com/onosproject/analytics/internal/channels"
	"github.com/segmentio/kafka-go"
)

func StartTopicReader(ctx context.Context, brokerURLs []string, topic string, groupID string) {

	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	output := channels.GetChannel(topic)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokerURLs,
		Topic:   topic,
		GroupID: groupID,
	})
	log.Printf("Starting kafka conection for %s for group id %s\n", topic, groupID)
	for {
		// the `ReadMessage` method blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}
		// after receiving the message, log its value
		*output <- string(msg.Value)
		log.Println("received: ", string(msg.Value))
	}
}
