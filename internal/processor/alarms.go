package processor

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"../channels"
	"github.com/segmentio/kafka-go"
)

func StartAlarmProcessor(channelName string, kafkaURI []string, kafkaTopic string) {
	input := channels.GetChannel(channelName)
	writer := kafkaWriter.GetWriter(kafkaURI, kafkaTopic)
	for {
		alarmJSON := <-input
		log.Printf("AlarmProcessor received %s", string(alarmJSON))
		var alarm models.Alarm
		if alarmJSON == "" {
			continue
		}
		err := json.Unmarshal([]byte(alarmJSON), &alarm)
		if err != nil {
			log.Fatalf("json.Unmarshal(%s) failed with %v", alarmJSON, err)
			break
		}
		enrichAlarm(&alarm)
		log.Printf("Alarm :%v", alarm)
		message, err := json.Marshal(alarm)
		if err != nil {
			log.Printf("failed to marshal alarm %v", err)
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
func enrichAlarm(alarm *models.Alarm) {
	alarm.Timestamp = time.Now()
}
