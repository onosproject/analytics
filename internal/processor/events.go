package processor

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"../channels"
	"github.com/onosproject/analytics/pkg/messages"
	"github.com/segmentio/kafka-go"
)

func StartEventProcessor(channelName string, kafkaURI []string, kafkaTopic string) {
	input := channels.GetChannel(channelName)
	producer := kafka.Writer{
		Addr:     kafka.TCP(kafkaURI[0]),
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
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
		msg := kafka.Message{Value: message}
		err = producer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Printf("producer.WriteMessages threw %v ", err)
		} else {
			log.Printf("Wrote messages to %s successfully", kafkaTopic)
		}
	}

}
func enrichEvent(event *messages.Event) {
	event.Timestamp = time.Now()
}
