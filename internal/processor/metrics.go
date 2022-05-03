package processor

import (
	"encoding/json"
	"time"

	"github.com/onosproject/analytics/pkg/messages"
)

func processMetric(metricJSON string) ([]byte, error) {
	var metric messages.Metric
	err := json.Unmarshal([]byte(metricJSON), &metric)
	if err != nil {
		return []byte{}, err
	}
	enrichMetric(&metric)
	log.Debugf("Metric after enrich: %v", metric)
	message, err := json.Marshal(metric)
	return message, err
}
func enrichMetric(metric *messages.Metric) {
	metric.Timestamp = time.Now()
}
