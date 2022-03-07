package listener

import (
	"context"

	"github.com/onosproject/analytics/internal/channels"
	"github.com/onosproject/analytics/pkg/kafkaClient"
)

func StartTopicReader(ctx context.Context, channelName string, brokerURLs []string, inbound string, groupID string) {

	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages
	output := channels.GetChannel(channelName)
	//blocks until shutdown
	kafkaClient.StartTopicReader(ctx, output, brokerURLs, inbound, groupID)

}
