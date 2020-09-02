package main

import (
	"log"

	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/GoEndpointManager/GoEndpointBackendManager"
)

//host = 10.60.1.20
//port = 18408
func TestSV() {
	svClient := StringBigsetService.NewStringBigsetServiceModel("", []string{}, GoEndpointBackendManager.EndPoint{
		Host:      "10.110.1.21",
		Port:      "185997",
		ServiceID: ""})
	bskey := generic.TStringKey("GiacNgoTVLinkPaper")
	// svClient.BsPutItem(bskey, &generic.TItem{
	// 	Key:   []byte("minhv2"),
	// 	Value: []byte("1234"),
	// })
	lsItems, err := svClient.BsGetSliceR(bskey, 0, 10)
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
