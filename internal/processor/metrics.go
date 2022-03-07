package processor

import (
	"encoding/json"
	"log"
	"time"

	"github.com/onosproject/analytics/internal/channels"
	"github.com/onosproject/analytics/pkg/kafkaClient"
	"github.com/onosproject/analytics/pkg/messages"
)

func StartMetricsProcessor(channelName string, kafkaURI []string, kafkaTopic string) {
	input := channels.GetChannel(channelName)

	writer := kafkaClient.GetWriter(kafkaURI[0], kafkaTopic)
	log.Printf("StartMetricsProcessor(%s,%s,%s)", channelName, kafkaURI[0], kafkaTopic)
	for {
		metricJSON := <-input
		log.Printf("MetricsProcessor received %s", string(metricJSON))
		var metric messages.Metric
		if metricJSON == "" {
			continue
		}
		err := json.Unmarshal([]byte(metricJSON), &metric)
		if err != nil {
			log.Fatalf("json.Unmarshal(%s) failed with %v", metricJSON, err)
			break
		}
		enrichMetric(&metric)
		log.Printf("Metric :%v", metric)
		message, err := json.Marshal(metric)
		if err != nil {
			log.Printf("failed to marshal metric %v", err)
			break
		}
		writer.SendMessage(message)
	}

}
func enrichMetric(metric *messages.Metric) {
	metric.Timestamp = time.Now()
}
