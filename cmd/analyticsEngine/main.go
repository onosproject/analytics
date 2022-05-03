/*
 * Copyright 2020-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */
package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"strings"
	"time"

	"github.com/onosproject/analytics/internal/configuration"
	"github.com/onosproject/analytics/internal/listener"
	"github.com/onosproject/onos-lib-go/pkg/logging"
)

var Config configuration.Configuration

func main() {
	var configFile = flag.String("conf", "analytics.json", "json file containing configuration")
	var logLevel = flag.String("logLevel", "error", "log level {error,warn,info,debug}")
	flag.Parse()
	var level logging.Level
	switch strings.ToUpper(*logLevel) {
	case "DEBUG":
		level = logging.DebugLevel
	case "INFO":
		level = logging.InfoLevel
	case "WARN":
		level = logging.WarnLevel
	case "ERROR":
		level = logging.ErrorLevel
	default:
		level = logging.ErrorLevel
	}
	log := logging.GetLogger("analytics")
	log.SetLevel(level)

	log.Infof("AnalyticsEngine Config: %s", *configFile)

	content, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatalf("Unable to load configuration file %v", err)
	}
	Config, err := configuration.GetConfiguration(content)
	if err != nil {
		log.Fatalf("Unable to load configuration %v", err)
	}
	js, _ := json.Marshal(Config)
	log.Debug(string(js))

	log.Debugf("Configuration: %v", Config)
	ctx := context.Background()

	topics := Config.Topics
	for i := 0; i < len(topics); i++ {
		for j := 0; j < len(topics[i].Brokers); j++ {
			var brokerURLs []string
			brokerURLs = append(brokerURLs, topics[i].Brokers[j].URL)
			log.Infof("calling listener.StartTopicReader(%v,%s,%v,%s,%s)",
				ctx, topics[i].Name, brokerURLs, topics[i].Queues.Inbound, Config.GroupID)
			go listener.StartTopicReader(ctx, topics[i].Name, brokerURLs, topics[i].Queues.Inbound, topics[i].Queues.Outbound, Config.GroupID)
		}
	}
	for {
		time.Sleep(time.Second * 60)
		log.Debug("AnalyticsEngine waking up")
		//TODO print stats etc
	}
}
