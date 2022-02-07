package channels
import (
    "fmt"
    )

var channelMap map[string]*chan string
func Init(){
  channelMap = make(map[string] *chan string)
}

func AddChannel(topicName string){
  fmt.Printf("Adding channel %s\n",topicName)
  channel := make(chan string)
  channelMap[topicName] = &channel
}

func GetChannel(topicName string)*chan string{
  fmt.Printf("Requesting channel %s\n",topicName)
  channel := channelMap[topicName]
  return channel
}

  
  
