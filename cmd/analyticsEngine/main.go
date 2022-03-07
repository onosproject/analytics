package main

import (
	"context"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/onosproject/analytics/internal/channels"
	"github.com/onosproject/analytics/internal/configuration"
	"github.com/onosproject/analytics/internal/listener"
	"github.com/onosproject/analytics/internal/processor"
)

var Config configuration.Configuration

func main() {
	var configFile = flag.String("conf", "analytics.json", "json file containing configuration")
	flag.Parse()
	log.Println(*configFile)

	content, err := ioutil.ReadFile(*configFile)
	if err != nil {
		log.Fatal(err)
	}
	Config, err := configuration.GetConfiguration(content)
	if err != nil {
		log.Fatalf("Unable to load configuration %v", err)
	}
	js, _ := json.Marshal(Config)
	log.Println(string(js))

	log.Println(Config)
	channels.Init()
	ctx := context.Background()

	topics := Config.Topics
	for i := 0; i < len(topics); i++ {
		channels.AddChannel(topics[i].Name)
		for j := 0; j < len(topics[i].Brokers); j++ {
			var brokerURLs []string
			brokerURLs = append(brokerURLs, topics[i].Brokers[j].URL)
			go listener.StartTopicReader(ctx, topics[i].Name, brokerURLs, topics[i].Queues.Inbound, Config.GroupID)
			switch strings.ToUpper(topics[i].Name) {
			case "ALARMS":
				go processor.StartAlarmProcessor(topics[i].Name, brokerURLs, topics[i].Queues.Outbound)
			case "EVENTS":
				go processor.StartEventProcessor(topics[i].Name, brokerURLs, topics[i].Queues.Outbound)
			case "METRICS":
				go processor.StartMetricsProcessor(topics[i].Name, brokerURLs, topics[i].Queues.Outbound)
			}
		}
	}
	for {
		time.Sleep(time.Second * 60)
		log.Println("AnalyticEngine Stats")
		channels.PrintChannelStats()
	}

}
