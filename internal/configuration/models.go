package configuration

import (
    "fmt"
    "encoding/json"
    )

type Broker struct{
  URL string `json:"url"`
}

type Topic struct {
  Name string `json:"name"`
  Brokers []Broker `json:"brokers"`
}

type Configuration struct {
  GroupID string `json:"group_id"`
  Topics []Topic `json:"topics"`
}

func GetConfiguration(config []byte) Configuration {
  var conf Configuration
  err := json.Unmarshal(config,&conf)
  if err != nil{
    fmt.Printf("Unable to unmarshal config file  %v",err)
  }
  return conf
}



