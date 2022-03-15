package channels

import (
	"github.com/onosproject/analytics/pkg/logger"
)

const CHANNEL_BUFFERSIZE = 20

var messageChannelMap map[string]chan string
var errorChannelMap map[string]chan error

func Init() {
	messageChannelMap = make(map[string]chan string)
	errorChannelMap = make(map[string]chan error)
}

func AddChannels(topicName string) (chan string, chan error) {
	if logger.IfDebug() {
		logger.Debug("Adding channels %s\n", topicName)
	}
	messageChannel := make(chan string, CHANNEL_BUFFERSIZE)
	messageChannelMap[topicName] = messageChannel

	errorChannel := make(chan error, CHANNEL_BUFFERSIZE)
	errorChannelMap[topicName] = errorChannel

	return messageChannel, errorChannel
}

func GetChannels(topicName string) (chan string, chan error) {
	if logger.IfDebug() {
		logger.Debug("Requesting channel %s\n", topicName)
	}
	return messageChannelMap[topicName],
		errorChannelMap[topicName]
}
func PrintChannelStats() {
	for channelName, channel := range messageChannelMap {
		logger.Info("Channel: %s, fill: %d", channelName, len(channel))
	}
}
