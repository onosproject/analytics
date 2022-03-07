package processor

import (
	"encoding/json"
	"log"
	"time"

	"github.com/onosproject/analytics/internal/channels"
	"github.com/onosproject/analytics/pkg/kafkaClient"
	"github.com/onosproject/analytics/pkg/messages"
)

func StartAlarmProcessor(channelName string, kafkaURI []string, kafkaTopic string) {
	input := channels.GetChannel(channelName)

	writer := kafkaClient.GetWriter(kafkaURI[0], kafkaTopic)
	for {
		alarmJSON := <-input
		log.Printf("AlarmProcessor received %s", string(alarmJSON))
		var alarm messages.Alarm
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
		writer.SendMessage(message)
	}

}
func enrichAlarm(alarm *messages.Alarm) {
	alarm.Timestamp = time.Now()
}
