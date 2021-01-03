package main

import (
	"log"

	"github.com/OpenStars/BackendService/StringBigsetService"
	"github.com/OpenStars/BackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
)

//host = 10.60.1.20
//port = 18408
func TestSV() {
	// 10.110.1.21:29810
	svClient := StringBigsetService.NewClient(nil, "/test/", "10.110.1.21", "29810")
	bskey := generic.TStringKey("GiacNgoTVLinkPaper")
	// svClient.BsPutItem(bskey, &generic.TItem{
	// 	Key:   []byte("minhv2"),
	// 	Value: []byte("1234"),
	// })
	lsItems, err := svClient.BsGetSlice(string(bskey), 0, 10)
	if err != nil {
		log.Println("[ERROR] err", err)
	}
	chanerr := make(chan error)
	<-chanerr
	if lsItems != nil {
		for i := 0; i < len(lsItems); i++ {
			log.Println(i, string(lsItems[i].Value), "key", string(lsItems[i].Key))
		}
	}

}
func main() {
	TestSV()
}
