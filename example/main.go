package main

import (
	"fmt"
	"go-event-delegation/broker"
)

func main() {
	brk := broker.New()
	subChan := make(chan []byte)
	brk.SubscribePlayerReadyEventV1(subChan)
	brk.PublishPlayerReadyEventV1(1)
	b := <-subChan
	fmt.Println(fmt.Sprintf("Player #%v is ready", b[0]))
}
