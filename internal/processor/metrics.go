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

func StartMetricsProcessor(channelName string, kafkaURI []string, kafkaTopic string) {
	input := channels.GetChannel(channelName)
	producer := kafka.Writer{
		Addr:     kafka.TCP(kafkaURI[0]),
		Topic:    kafkaTopic,
		Balancer: &kafka.LeastBytes{},
	}
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
		msg := kafka.Message{Value: message}
		err = producer.WriteMessages(context.Background(), msg)
		if err != nil {
			log.Printf("producer.WriteMessages threw %v ", err)
		} else {
			log.Printf("Wrote messages to %s successfully", kafkaTopic)
		}
	}

}
func enrichMetric(metric *messages.Metric) {
	metric.Timestamp = time.Now()
}
