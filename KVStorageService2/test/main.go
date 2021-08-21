package main

import (
	"log"

	KVStorageService "github.com/OpenStars/BackendService/KVStorageService2"
)

func main() {
	kvstorageclient := KVStorageService.NewClient("/tets", "127.0.0.1", "8883")
	kvstorageclient.PutData("1", []byte("lehaison"))
	item, err := kvstorageclient.GetData("1")
	log.Println("multiget", string(item), "err", err)
	// for i := 0; i < 1000; i++ {
	// 	kvstorageclient.PutData(strconv.Itoa(i), strconv.Itoa(i*i))
	// }
	// for i := 0; i < 1000; i++ {
	// 	value, err := kvstorageclient.GetData(strconv.Itoa(i))
	// 	log.Println("key", i, "value", value, "err", err)
	// }
	// kvstorageclient.RemoveData("7")
	// kvstorageclient.RemoveData("15")
	// kvstorageclient.RemoveData("20")
	// results, missingkey, err := kvstorageclient.GetListData([]string{"1", "3", "7", "20", "21", "8"})
	// log.Println("result", results, "missingkey", missingkey, "err", err)
}
