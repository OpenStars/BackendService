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
	svClient := StringBigsetService.NewClient(nil, "/test/", "10.9.0.17", "18407")
	rfailed, err := svClient.BsMultiRemoveBsItem([]*generic.TBigsetItem{
		&generic.TBigsetItem{
			Bskey:   []byte("bskey1"),
			Itemkey: []byte("itemkey1"),
		},
		&generic.TBigsetItem{
			Bskey:   []byte("bskey3"),
			Itemkey: []byte("itemkey2"),
		},
	})
	log.Println("total failed ", len(rfailed))
	log.Println("rfailed", string(rfailed[0].Bskey), "err", err)
	// r, err := svClient.BsMultiPutBsItem([]*generic.TBigsetItem{
	// 	&generic.TBigsetItem{
	// 		Bskey:     []byte("bskey1"),
	// 		Itemkey:   []byte("itemkey1"),
	// 		Itemvalue: []byte("itemval1"),
	// 	},
	// 	&generic.TBigsetItem{
	// 		Bskey:     []byte("bskey1"),
	// 		Itemkey:   []byte("itemkey2"),
	// 		Itemvalue: []byte("itemval2"),
	// 	},
	// 	&generic.TBigsetItem{
	// 		Bskey:     []byte("bskey2"),
	// 		Itemkey:   []byte("itemkey1"),
	// 		Itemvalue: []byte("itemval1"),
	// 	},
	// })
	// item, _ := svClient.BsGetItem("bskey1", "itemkey1")
	// log.Println("item value", string(item.Value))
	// bskey := generic.TStringKey("GiacNgoTVLinkPaper")
	// // svClient.BsPutItem(bskey, &generic.TItem{
	// // 	Key:   []byte("minhv2"),
	// // 	Value: []byte("1234"),
	// // })
	// lsItems, err := svClient.BsGetSlice(string(bskey), 0, 10)
	// if err != nil {
	// 	log.Println("[ERROR] err", err)
	// }
	// chanerr := make(chan error)
	// <-chanerr
	// if lsItems != nil {
	// 	for i := 0; i < len(lsItems); i++ {
	// 		log.Println(i, string(lsItems[i].Value), "key", string(lsItems[i].Key))
	// 	}
	// }

}
func main() {
	TestSV()
	// databigset := StringBigsetService.NewClient(nil, "/test", "10.110.69.96", "20547")
	// if databigset == nil {
	// 	return
	// }
	// log.Println("connect success")
	// ok, err := databigset.BsPutItem("PRAY_INFO", utils.PaddingZeros(3467), utils.PaddingZeros(3487))
	// log.Println("ok", ok, "err", err)
	// item, err := databigset.BsGetItem("PRAY_INFO", utils.PaddingZeros(3467))
	// if err != nil {
	// 	log.Println("err", item)
	// }
	// if item != nil {
	// 	log.Println("item", string(item.Value))
	// }
	// log.Println("aaaa")
}
