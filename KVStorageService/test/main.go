package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/OpenStars/BackendService/KVStorageService"
)

func main() {
	kvstorageclient := KVStorageService.NewClient("/tets", "10.120.1.11", "18765")

	// kvstorageclient.PutData("sh1", "le hai son")
	// kvstorageclient.PutData("sh2", "tran minh tuan")
	// kvstorageclient.PutData("sh3", "nguyen thi mai anh")
	// kvstorageclient.PutData("sh4", "tran van lam")
	for j := 0; j < 100; j++ {
		// fmt.Print(j)
		// fmt.Println(kvstorageclient.PutData("sh1", "lhs"))
		go func(i int) {
			fmt.Print(i)
			fmt.Println(kvstorageclient.PutData("sh1", "lhs"))
		}(j)
	}

	// }
	reader := bufio.NewReader(os.Stdin)
	for {
		reader.ReadString('\n')
		for i := 0; i < 2; i++ {
			go func() {
				fmt.Println("put data")
				fmt.Println(kvstorageclient.PutData("sh1", "lhs"))
			}()
		}

	}
	waitKEy := make(chan bool)
	<-waitKEy
	// sessionKey, err := kvstorageclient.OpenIterate()
	// if err != nil {
	// 	log.Fatalln("open iterate err", err)
	// }
	// fmt.Println("sessionkey", sessionKey)
	// defer kvstorageclient.CloseIterate(sessionKey)
	// for {
	// 	lsItem, err := kvstorageclient.NextListItems(sessionKey, 10)
	// 	if err != nil {
	// 		log.Fatalln("err data", err)
	// 	}
	// 	for _, item := range lsItem {
	// 		fmt.Println("key", item.Key, "value", item.Value)
	// 	}

	// }

}
