package main

import (
    "fmt"
    "flag"
    "io/ioutil"
    "log"
    "context"
    "time"
  //  "encoding/json"

    "github.com/onosproject/analytics/internal/configuration"
    "github.com/onosproject/analytics/internal/channels"
    "github.com/onosproject/analytics/internal/kafkaConnector"
)

var Config configuration.Configuration

func main() {
  var configFile = flag.String("conf","analytics.json","json file containing configuration")
  flag.Parse()
  fmt.Println(*configFile)

  content, err := ioutil.ReadFile(*configFile)
    if err != nil {
        log.Fatal(err)
  }
  Config = configuration.GetConfiguration(content)
  fmt.Println(Config)
  ctx := context.Background()
  channels.Init()


  topics:= Config.Topics
  for i:=0;i<len(topics);i++{
    channels.AddChannel(topics[i].Name)
    for j:=0; j<len(topics[i].Brokers);j++{
    var brokerURLs []string
      brokerURLs=append(brokerURLs,topics[i].Brokers[j].URL)
      go kafkaConnector.StartTopicReader(ctx, brokerURLs, topics[i].Name, Config.GroupID)
    }
  }
  time.Sleep(time.Second * 60)
}
