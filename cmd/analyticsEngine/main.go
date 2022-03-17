
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
	"github.com/onosproject/analytics/pkg/logger"
)

var Config configuration.Configuration

func main() {
	var configFile = flag.String("conf", "analytics.json", "json file containing configuration")
	var logFile = flag.String("logFile", "AnalyticEngine.log", "file name to log to")
	var logLevel = flag.String("logLevel", "error", "log level {error,warn,info,debug}")
	flag.Parse()
	var level logger.LogLevel
	switch strings.ToUpper(*logLevel) {
	case "DEBUG":
		level = logger.DEBUG
	case "INFO":
		level = logger.INFO
	case "WARN":
		level = logger.WARN
	case "ERROR":
		level = logger.ERROR
	default:
		level = logger.ERROR
	}

	logger.Init(*logFile, level)
	if logger.IfInfo() {
		logger.Info("AnalyticsEngine Config: %s", *configFile)
	}

	content, err := ioutil.ReadFile(*configFile)
	if err != nil {
		logger.Fatal("Unable to load configuration file %v", err)
	}
	Config, err := configuration.GetConfiguration(content)
	if err != nil {
		logger.Fatal("Unable to load configuration %v", err)
	}
	if logger.IfDebug() {
		js, _ := json.Marshal(Config)
		logger.Debug(string(js))

		logger.Debug("Configuration: %v", Config)
	}
	ctx := context.Background()

	topics := Config.Topics
	for i := 0; i < len(topics); i++ {
		for j := 0; j < len(topics[i].Brokers); j++ {
			var brokerURLs []string
			brokerURLs = append(brokerURLs, topics[i].Brokers[j].URL)
			if logger.IfInfo() {
				logger.Info("calling listener.StartTopicReader(%v,%s,%v,%s,%s)",
					ctx, topics[i].Name, brokerURLs, topics[i].Queues.Inbound, Config.GroupID)

			}
			go listener.StartTopicReader(ctx, topics[i].Name, brokerURLs, topics[i].Queues.Inbound, topics[i].Queues.Outbound, Config.GroupID)
		}
	}
	for {
		time.Sleep(time.Second * 60)
		if logger.IfDebug() {
			logger.Debug("AnalyticsEngine waking up")
		}
		//TODO print stats etc
	}
}
