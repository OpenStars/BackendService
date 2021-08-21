package main

import (
	"fmt"

	"github.com/OpenStars/BackendService/KVStorageService"
)

func main() {
	kvstorageclient := KVStorageService.NewClient("/tets", "10.60.68.100", "8873")
	fmt.Println(kvstorageclient.NextItem(1))
	// fmt.Println(kvstorageclient.OpenIterate())
	// kvstorageclient.PutData("le", "a")
	// kvstorageclient.PutData("hai", "b")
	// kvstorageclient.PutData("son", "c")
	// kvstorageclient.PutData("4", "d")
	// fmt.Println(kvstorageclient.GetData("1"))
	// fmt.Println(kvstorageclient.GetData("2"))
	// fmt.Println(kvstorageclient.GetData("3"))
	// fmt.Println(kvstorageclient.GetData("4"))
	// err := kvstorageclient.OpenIterate()
	// fmt.Println(err)
	// item, err := kvstorageclient.NextItem()
	// if err != nil {
	// 	fmt.Println("err", err)
	// 	return
	// }
	// fmt.Println("key", item.Key, "value", item.Value)
	// kvstorageclient.PutData("2", "2")
	// item, err := kvstorageclient.GetData("2")
	// log.Println("multiget", item, "err", err)
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
