package kafkaClient

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func StartTopicReader(ctx context.Context, channel chan string, brokerURLs []string, inbound string, groupID string) {

	// initialize a new reader with the brokers and topic
	// the groupID identifies the consumer and prevents
	// it from receiving duplicate messages

	brokerStr := "Brokers: "
	for i := 0; i < len(brokerURLs); i++ {
		brokerStr += brokerURLs[i]
	}
	log.Printf("StartTopicReader(%s,%s,%s)", brokerStr, inbound, groupID)
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: brokerURLs,
		Topic:   inbound,
		GroupID: groupID,
	})
	log.Printf("Starting kafka conection for %s for group id %s\n", inbound, groupID)
	for {
		// the `ReadMessage` method blocks until we receive the next event
		log.Println("Before ReadMessage")
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			//TODO don't panic in production
			panic("could not read message " + err.Error())
		}
		log.Println("received: ", string(msg.Value))

		channel <- string(msg.Value)
		log.Println("sent: ", string(msg.Value))
	}

}
