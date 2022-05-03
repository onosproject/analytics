/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package configuration

import (
	"encoding/json"

	"github.com/onosproject/onos-lib-go/pkg/logging"
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

var log = logging.GetLogger("configuration")

/*
GetConfiguration
converts byte array from config file into a configuration
struct
*/
func GetConfiguration(config []byte) (Configuration, error) {
	var conf Configuration
	err := json.Unmarshal(config, &conf)
	if err != nil {
		log.Errorf("Unable to unmarshal config file  %v", err)
		return conf, err
	}
	log.Debugf("Created Configuration Object %v", conf)
	return conf, nil
}
