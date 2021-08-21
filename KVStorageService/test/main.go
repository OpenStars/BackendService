package main

import (
	"fmt"
	"log"

	"github.com/OpenStars/BackendService/KVStorageService"
)

func main() {
	kvstorageclient := KVStorageService.NewClient("/tets", "10.60.68.100", "8873")
	sessionKey, err := kvstorageclient.OpenIterate()
	if err != nil {
		log.Fatalln("open iterate err", err)
	}
	defer kvstorageclient.CloseIterate(sessionKey)
	for {
		item, err := kvstorageclient.NextItem(sessionKey)
		if err != nil {
			break
		}
		fmt.Println("key", item.Key, "value", item.Value)
	}

}
