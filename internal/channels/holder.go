/*
 * Copyright 2022-present Open Networking Foundation
 *
 * SPDX-License-Identifier: Apache-2.0
 *
 */

package channels

import (
	"log"
)

var channelMap map[string]*chan string

func Init() {
	channelMap = make(map[string]*chan string)
}

func AddChannel(topicName string) {
	log.Printf("Adding channel %s\n", topicName)
	channel := make(chan string)
	channelMap[topicName] = &channel
}

func GetChannel(topicName string) *chan string {
	log.Printf("Requesting channel %s\n", topicName)
	channel := channelMap[topicName]
	return channel
}
