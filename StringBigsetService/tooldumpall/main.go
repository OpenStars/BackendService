package main

import (
	"fmt"

	"github.com/OpenStars/BackendService/StringBigsetService"
)

func main() {
	bigset := StringBigsetService.NewClient(nil, "/test/bigsetfortest", "10.110.69.96", "20507")
	totalBsKey, err := bigset.TotalStringKeyCount()
	fmt.Println("total bigset key", totalBsKey, err)
	// offset := 0
	// count := 1000
	// for {
	// 	lsKey, err := bigset.GetListKey(int64(offset), int32(count))
	// 	if err != nil {
	// 		log.Println("lsKey get error: ", err)
	// 		break
	// 	}
	// 	for _, bskey := range lsKey {
	// 		fmt.Println("bskey", bskey)
	// 	}
	// 	offset += count
	// }
}
