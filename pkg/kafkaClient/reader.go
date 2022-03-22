package kafkaClient

import (
	"context"

	"github.com/onosproject/analytics/pkg/logger"
	"github.com/segmentio/kafka-go"
)

/*
StartTopicReader
Connects to kafka topic and reads new messages and writes them to the messageChannel any errors will be sent to errorChannel
*/
func StartTopicReader(ctx context.Context, messageChannel chan string, errorChannel chan error, brokerURLs []string, inbound string, groupID string) {
	if logger.IfDebug() {
		allBrokers := ""
		for _, broker := range brokerURLs {
			allBrokers += broker
			allBrokers += ","
		}
		logger.Debug("StartTopicReader(%s,%s,%s)", allBrokers, inbound, groupID)
	}

	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokerURLs,
		Topic:   inbound,
		GroupID: groupID,
	})

	brokerStr := "Brokers: "
	for i := 0; i < len(brokerURLs); i++ {
		brokerStr += brokerURLs[i]
	}
	if logger.IfInfo() {
		logger.Info("StartTopicReader(%s,%s,%s)", brokerStr, inbound, groupID)
	}
	for {
		// the `ReadMessage` function blocks until we receive the next event
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			logger.Error("Error reading off kafka bus err:%v", err)
			errorChannel <- err
			continue
		}
		messageChannel <- string(msg.Value)
	}
}
