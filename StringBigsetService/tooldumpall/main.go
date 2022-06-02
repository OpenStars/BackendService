package main

import (
	"fmt"
	"log"

	"github.com/OpenStars/BackendService/StringBigsetService"
)

func main() {
	bigset := StringBigsetService.NewClient(nil, "/test/bigsetfortest", "10.60.1.20", "18407")
	totalBsKey, err := bigset.TotalStringKeyCount()
	fmt.Println("total bigset key", totalBsKey, err)
	offset := 0
	count := 1000
	for {
		lsKey, err := bigset.GetListKey(int64(offset), int32(count))
		if err != nil {
			log.Println("lsKey get error: ", err)
			break
		}
		if len(lsKey) == 0 {
			break
		}
		for _, bskey := range lsKey {
			fmt.Println("bskey", bskey)
			fmt.Println("lst")
		}
		offset += count
	}
	log.Println("done")
}
