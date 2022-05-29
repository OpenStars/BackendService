package main

import (
	"fmt"
	"log"

	"github.com/OpenStars/BackendService/StringBigsetService"
)

var bigset StringBigsetService.Client

func main() {
	bigset = StringBigsetService.NewClient(nil, "/test/checkbigset", "10.60.68.102", "22407")
	if bigset == nil {
		log.Fatalln("Cannot connect to bigset service")
	}
	offset := 0
	count := 1000
	totalBsKey, err := bigset.TotalStringKeyCount()
	totalBsKeyCounter := 0
	fmt.Println("totalBigsetKey", totalBsKey, err)
	for {
		fmt.Println("[INFO] offset bskey", offset)
		lsKey, err := bigset.GetListKey(int64(offset), int32(count))
		if err != nil {
			log.Println("lsKey get error: ", err)
			break
		}
		if len(lsKey) == 0 {
			break
		}
		totalBsKeyCounter += len(lsKey)
		for _, bskey := range lsKey {
			totalItemKey, _ := bigset.GetTotalCount(bskey)
			totalItemKeyCount := 0
			offsetItemKey := 0
			for {
				lsItem, _ := bigset.BsGetSlice(bskey, int32(offsetItemKey), int32(count))
				if len(lsItem) == 0 {
					break
				}
				currentItemKey := string(lsItem[0].Key)
				for _, item := range lsItem {
					if string(currentItemKey) > string(item.Key) {
						log.Println("[DATA ERROR] key sort failed", "currentItemKey", string(currentItemKey), "nextItemKey", string(item.Key), "bigsetkey", bskey)
					}
					currentItemKey = string(item.Key)
				}
				totalItemKeyCount += len(lsItem)
				offsetItemKey += count
			}
			if totalItemKeyCount != int(totalItemKey) {
				log.Println("[DATA ERROR] total item key not equal to total item key count: ", totalItemKey, "vs", totalItemKeyCount, "bskey", bskey)
			}

		}
		offset += count
	}
	if totalBsKey != int64(totalBsKeyCounter) {
		log.Println("[DATA ERROR] total bigset key not equal to bigset key count ", totalBsKey, "vs", totalBsKeyCounter)
	}
	fmt.Println("[INFO] check done")
}
