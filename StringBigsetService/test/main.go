package main

import (
	"log"

	"github.com/OpenStars/BackendService/StringBigsetService"
	"github.com/OpenStars/BackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"gitlab.123xe.vn/TrustKeysV2/socialnetworks/Like/utils"
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
	databigset := StringBigsetService.NewClient(nil, "/test", "10.110.69.96", "20547")
	ok, err := databigset.BsPutItem("PRAY_INFO", utils.PaddingZeros(3467), utils.PaddingZeros(3487))
	log.Println("ok", ok, "err", err)
	// item, err := databigset.BsGetItem("PRAY_INFO", utils.PaddingZeros(3467))
	// if err != nil {
	// 	log.Println("err", item)
	// }
	// if item != nil {
	// 	log.Println("item", string(item.Value))
	// }
}
