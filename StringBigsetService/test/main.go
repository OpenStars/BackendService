package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/OpenStars/BackendService/StringBigsetService"
)

var bigset StringBigsetService.Client

func GetItem() {
	// |PostUIDFollower|0000000000000793479
	// |TimeFollowUIDPrefix|0000000000000000005
	lsItems, err := bigset.BsGetSlice("this_is_bigset_key_of_my_bigset_from_os_linux_corei5_ram8gb_ssd256_id_100", 1472436, 20000)
	if err != nil {
		log.Fatalln(err)
	}
	total, _ := bigset.GetTotalCount("this_is_bigset_key_of_my_bigset_from_os_linux_corei5_ram8gb_ssd256_id_100")
	fmt.Println("max size item", len(lsItems), "total", total)
	// fmt.Println(lsItems)
}
func ListAllItem() {
	startIndex := int64(7000)
	numItem := int32(1000)
	for {
		lsKey, err := bigset.GetListKey(startIndex, numItem)
		if err != nil || len(lsKey) == 0 {
			log.Fatalln("get list key", err)
		}
		for i, bskey := range lsKey {
			totalItem, err := bigset.GetTotalCount(bskey)
			if err != nil {
				log.Fatalln(err)
			}

			startIndexItem := int32(0)
			totalRealItem := 0
			for startIndexItem < int32(totalItem) {
				lsItems, err := bigset.BsGetSlice(bskey, startIndexItem, numItem)
				if err != nil {
					log.Fatalln("item index", startIndexItem, err)
				}
				startIndexItem += numItem
				totalRealItem += len(lsItems)
			}
			if totalItem > 0 {
				log.Println(startIndex+int64(i), "bskey", bskey, "total Item", totalItem, "travel item", totalRealItem)
			}
		}
		startIndex += int64(numItem)
	}
}

func bulkPut() {
	for bskey := 0; bskey < 100; bskey++ {
		go func(key int) {
			for i := 0; i < 10000000; i++ {
				b := rand.Int()
				bskey := fmt.Sprintf("this_is_bigset_key_of_my_bigset_from_os_linux_corei5_ram8gb_ssd256_id_%d", bskey)
				itemkey := fmt.Sprintf("%19d", b)
				fmt.Println("bskey", bskey, "itemkey", itemkey)
				bigset.BsPutItem(bskey, itemkey, itemkey)
			}
		}(bskey)
	}
	waitKey := make(chan bool)
	<-waitKey
}
func main() {

	bigset = StringBigsetService.NewClient(nil, "/test", "10.110.69.96", "20527")
	// lsItems, err := bigset.BsRangeQuery("eth:NormalAddress", "0", "zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println("len item", len(lsItems))
	// listKey, err := bigset.GetListKey(0, 1000)
	// fmt.Println("listkey", len(listKey), err)
	// fmt.Println(listKey)
	// GetItem()
	ListAllItem()
	// for bskey := 0; bskey < 100; bskey++ {
	// 	go func(key int) {
	// 		for i := 0; i < 10000000; i++ {
	// 			b := rand.Int()
	// 			bskey := fmt.Sprintf("this_is_bigset_key_of_my_bigset_from_os_linux_corei5_ram8gb_ssd256_id_%d", bskey)
	// 			itemkey := fmt.Sprintf("%19d", b)
	// 			fmt.Println("bskey", bskey, "itemkey", itemkey)
	// 			bigset.BsPutItem(bskey, itemkey, itemkey)
	// 		}
	// 	}(bskey)
	// }
	// waitKey := make(chan bool)
	// <-waitKey
	// r, err := bigset.BsRangeQueryAll("eth:NormalAddress")
	// fmt.Println("r", r, "err", err)
	// GetItem()
	// ListAllItem()
	// bigset.BsRangeQuery("eth:NormalAddress", "", "")
	// totalCount, err := bigset.TotalStringKeyCount()
	// fmt.Println("1", totalCount, err)
	// totalCount, err = bigset.TotalStringKeyCount()
	// fmt.Println("2", totalCount, err)
	// totalCount, err = bigset.TotalStringKeyCount()
	// fmt.Println("3", totalCount, err)
}
