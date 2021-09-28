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
	lsItems, err := bigset.BsGetSlice("this_is_bigset_key_of_my_bigset_from_os_linux_corei5_ram8gb_ssd256_id_100", 1200000, 1000)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("len items", len(lsItems))
	total, _ := bigset.GetTotalCount("this_is_bigset_key_of_my_bigset_from_os_linux_corei5_ram8gb_ssd256_id_100")
	fmt.Println("max size item total", total)
	// fmt.Println(lsItems)
}

func DeleteRandom() {
	totalItem, err := bigset.GetTotalCount("this_is_bigset_key_of_my_bigset_from_os_linux_corei5_ram8gb_ssd256_id_99")
	log.Println("total", totalItem, err)
	for i := 0; i < 1000; i++ {

		pos := rand.Int31n(10)
		item, err := bigset.BsGetSlice("", pos, 1)
		if err != nil || item == nil || len(item) == 0 {
			log.Println("[ERROR] get slice at", pos, err)
			continue
		}
		ok, err := bigset.BsRemoveItem("this_is_bigset_key_of_my_bigset_from_os_linux_corei5_ram8gb_ssd256_id_100", string(item[0].Key))
		if err != nil {
			log.Fatalln("delete", err)
		}
		if !ok {
			log.Println("[ERROR] bsremove at ", pos, err)
		}
		fmt.Println("remove item at", pos, "oke")

	}
}

func DeleteItem() {
	bigset.BsRemoveItem("this_is_bigset_key_of_my_bigset_from_os_linux_corei5_ram8gb_ssd256_id_100", "7957797661344512522")
}
func ListAllItem() {
	startIndex := int64(0)
	numItem := int32(1000)
	for {
		lsKey, err := bigset.GetListKey(startIndex, numItem)
		if err != nil || len(lsKey) == 0 {
			log.Fatalln("get list key", err)
		}
		for i, bskey := range lsKey {
			log.Println("bskey", bskey)
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
				// for _, item := range lsItems {
				// 	log.Println("delete item key", string(item.Key))
				// 	ok, err := bigset.BsRemoveItem(bskey, string(item.Key))
				// 	if ok {
				// 		log.Println("thang nay xoa duoc ne")
				// 	}
				// 	if err != nil {
				// 		log.Fatalln(err)
				// 	}
				// 	fmt.Println("delete", ok, err)

				// }
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
			for i := 0; i < 2000000; i++ {
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

	bigset = StringBigsetService.NewClient(nil, "/test", "10.110.69.96", "18407")
	// DeleteItem()
	// DeleteRandom()
	// log.Println("delete oke")
	// GetItem()
	// bulkPut()
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
	waitKey := make(chan bool)
	<-waitKey
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
