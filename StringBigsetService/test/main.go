package main

import (
	"fmt"

	"github.com/OpenStars/BackendService/StringBigsetService"
)

func main() {
	bigset := StringBigsetService.NewClient(nil, "/test/", "10.110.69.96", "22147")
	totalCount, err := bigset.TotalStringKeyCount()
	fmt.Println("1", totalCount, err)
	totalCount, err = bigset.TotalStringKeyCount()
	fmt.Println("2", totalCount, err)
	totalCount, err = bigset.TotalStringKeyCount()
	fmt.Println("3", totalCount, err)
}
