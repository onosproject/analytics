package channels

import (
	"log"
)

const CHANNEL_BUFFERSIZE = 20

var channelMap map[string]chan string

func Init() {
	channelMap = make(map[string]chan string)
}

func AddChannel(topicName string) {
	log.Printf("Adding channel %s\n", topicName)
	channel := make(chan string, CHANNEL_BUFFERSIZE)
	channelMap[topicName] = channel
}

func GetChannel(topicName string) chan string {
	log.Printf("Requesting channel %s\n", topicName)
	channel := channelMap[topicName]
	return channel
}
func PrintChannelStats(){
	for channelName, channel := range channelMap{
		log.Printf("Channel: %s, fill: %d",channelName, len(channel))
	}
}
