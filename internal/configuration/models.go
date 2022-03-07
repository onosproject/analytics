package configuration

import (
	"encoding/json"
	"log"
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
		log.Printf("Unable to unmarshal config file  %v", err)
		return conf, err
	}
	return conf, nil
}
