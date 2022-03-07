package listener

import (
	"context"

	"../../pkg/kafkaClient"
	"../channels"
)

func StartTopicReader(ctx context.Context, channelName string, brokerURLs []string, inbound string, groupID string) {

	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	output := channels.GetChannel(channelName)
	//blocks until shutdown
	kafkaClient.StartTopicReader(ctx, output, brokerURLs, inbound, groupID)

}
