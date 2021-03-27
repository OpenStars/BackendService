package main

import (
	"log"

	"github.com/OpenStars/BackendService/KVCounterService"
)

func main() {
	kvcounter := KVCounterService.NewClient(nil, "/test", "127.0.0.1", "12004")
	listkey := []string{"1", "2", "3"}
	listItems, err := kvcounter.GetMultiValue(listkey)
	if err != nil {
		log.Fatalln("err", err)
	}
	for _, item := range listItems {
		log.Println("key", item.Key, "value", item.Value)
	}
}
