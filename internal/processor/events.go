package processor

import (
	"encoding/json"
	"log"
	"time"

	"github.com/onosproject/analytics/internal/channels"
	"github.com/onosproject/analytics/pkg/kafkaClient"
	"github.com/onosproject/analytics/pkg/messages"
)

func StartEventProcessor(channelName string, kafkaURI []string, kafkaTopic string) {
	input := channels.GetChannel(channelName)
	writer := kafkaClient.GetWriter(kafkaURI[0], kafkaTopic)

	log.Printf("StartEventProcessor(%s,%s,%s)", channelName, kafkaURI[0], kafkaTopic)
	for {
		eventJSON := <-input
		log.Printf("EventProcessor received %s", string(eventJSON))
		var event messages.Event
		if eventJSON == "" {
			continue
		}
		err := json.Unmarshal([]byte(eventJSON), &event)
		if err != nil {
			log.Fatalf("json.Unmarshal(%s) failed with %v", eventJSON, err)
			break
		}
		enrichEvent(&event)
		log.Printf("Event :%v", event)
		message, err := json.Marshal(event)
		if err != nil {
			log.Printf("failed to marshal event %v", err)
			break
		}
		writer.SendMessage(message)
	}

}
func enrichEvent(event *messages.Event) {
	event.Timestamp = time.Now()
}
