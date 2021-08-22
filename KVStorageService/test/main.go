package main

import (
	"fmt"
	"log"

	"github.com/OpenStars/BackendService/KVStorageService"
)

func main() {
	kvstorageclient := KVStorageService.NewClient("/tets", "10.60.68.100", "8873")
	// kvstorageclient.PutData("sh1", "le hai son")
	// kvstorageclient.PutData("sh2", "tran minh tuan")
	// kvstorageclient.PutData("sh3", "nguyen thi mai anh")
	// kvstorageclient.PutData("sh4", "tran van lam")
	sessionKey, err := kvstorageclient.OpenIterate()
	if err != nil {
		log.Fatalln("open iterate err", err)
	}
	fmt.Println("sessionkey", sessionKey)
	defer kvstorageclient.CloseIterate(sessionKey)
	for {
		lsItem, err := kvstorageclient.NextListItems(sessionKey, 10)
		if err != nil {
			log.Fatalln("err data", err)
		}
		for _, item := range lsItem {
			fmt.Println("key", item.Key, "value", item.Value)
		}

	}

}
