package main

import (
	"fmt"

	"github.com/OpenStars/BackendService/KVCounterService"
)

func main() {
	kvcounter := KVCounterService.NewClient(nil, "/test", "10.60.68.103", "7974")

	k, err := kvcounter.GetValue("test")
	fmt.Println("k", k, "err", err)
	// listkey := []string{"1", "2", "3"}
	// listItems, err := kvcounter.GetMultiValue(listkey)
	// if err != nil {
	// 	log.Fatalln("err", err)
	// }
	// for _, item := range listItems {
	// 	log.Println("key", item.Key, "value", item.Value)
	// }
}
