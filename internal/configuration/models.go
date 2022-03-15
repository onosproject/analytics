package configuration

import (
	"encoding/json"

	"github.com/onosproject/analytics/pkg/logger"
)

type Queue struct {
	Inbound  string `json:"inbound"`
	Outbound string `json:"outbound"`
}
type Broker struct {
	URL string `json:"url"`
}

type Topic struct {
	Name    string   `json:"name"`
	Brokers []Broker `json:"brokers"`
	Queues  Queue    `json:"queues"`
}

type Configuration struct {
	GroupID string  `json:"group_id"`
	Topics  []Topic `json:"topics"`
}

func GetConfiguration(config []byte) (Configuration, error) {
	var conf Configuration
	err := json.Unmarshal(config, &conf)
	if err != nil {
		logger.Error("Unable to unmarshal config file  %v", err)
		return conf, err
	}
	if logger.IfDebug() {
		logger.Debug("Created Configuration Object %v", conf)
	}
	return conf, nil
}
