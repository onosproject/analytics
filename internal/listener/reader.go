package listener

import (
	"context"

	"github.com/onosproject/analytics/internal/processor"
	"github.com/onosproject/analytics/pkg/kafkaClient"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var log = logging.GetLogger("listener")

/*
StartTopicReader
creates needed channels and wires together the kafkaClient and processor
*/
func StartTopicReader(ctx context.Context, channelName string, brokerURLs []string,
	inbound string, outbound string, groupID string) {
	log.Infof("Calling processor.StartProcessor(%s,%v,%s)",
		channelName, brokerURLs, outbound)

	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	messageChan := make(chan string)
	errorChan := make(chan error)
	//blocks until shutdown
	go processor.StartProcessor(channelName, messageChan, errorChan, brokerURLs, outbound)
	kafkaClient.StartTopicReader(ctx, messageChan, errorChan, brokerURLs, inbound, groupID)
}
