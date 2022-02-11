/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package configuration

import (
	"encoding/json"
	"testing"
)

var configurationJson = `{
  "group_id": "cutty-sark",
  "topics": [
    {
      "name": "Alarms",
      "brokers": [
        {
          "url": "miner:9092"
        },
        {
          "url": "miner:9093"
        }
      ]
    },
    {
      "name": "Events",
      "brokers": [
        {
          "url": "miner:9092"
        },
        {
          "url": "miner:9093"
        }
      ]
    },
    {
      "name": "Metrics",
      "brokers": [
        {
          "url": "miner:9092"
        },
        {
          "url": "miner:9093"
        }
      ]
    }
  ]
}`
var badConfigurationJson = `{
  "groupId": "cutty-sark",
  "mytopicname": [
    {
      "name": "Alarms",
      "broke": [
        {
          "url": "miner:9092"
        },
        {
          "url": "miner:9093"
        }
      ]
    },
    {
      "name": "Events",
      "broke": [
        {
          "url": "miner:9092"
        },
        {
          "url": "miner:9093"
        }
      ]
    },
    {
      "name": "Metrics",
      "broke": [
        {
          "url": "miner:9092"
        },
        {
          "url": "miner:9093"
        }
      ]
    }
  ]
}`

func TestGetConfiguration(t *testing.T) {
	confObject, err := GetConfiguration([]byte(configurationJson))
	if err != nil {
		t.Error("Error marshalling configurationJson")
	}
	bytes, _ := json.MarshalIndent(confObject, "", "  ")

	match := compare(configurationJson, string(bytes))
	if !match {
		t.Error("configurationJson should not match")
	}

	confObject, _ = GetConfiguration([]byte(badConfigurationJson))
	bytes, _ = json.MarshalIndent(confObject, "", "  ")
	match = compare(configurationJson, string(bytes))
	if match {
		t.Error("badConfigurationJson should not match")
	}
}
func compare(original string, rendered string) bool {
	return original == rendered
}
