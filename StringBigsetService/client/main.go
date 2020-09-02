package main

import (
	"flag"
	"log"
	"strings"

	"github.com/OpenStars/BackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/EtcdBackendService/StringBigsetService"
)

func BsGetItem(host, port, bskey, itemkey string) {

	stringbs := StringBigsetService.NewStringBigsetServiceModel2(nil, "/test", host, port)
	item, err := stringbs.BsGetItem(generic.TStringKey(bskey), generic.TItemKey(itemkey))
	if err != nil || item == nil {
		log.Println("[ERROR] BsGetItem bskey", bskey, "itemkey", itemkey, "err", err)
	}
	log.Println("itemkey ", itemkey, "itemvalue", string(item.Value))
}

func BsPutItem(host, port, bskey, itemkey, itemvalue string) {
	stringbs := StringBigsetService.NewStringBigsetServiceModel2(nil, "/test", host, port)
	err := stringbs.BsPutItem(generic.TStringKey(bskey), &generic.TItem{
		Key:   []byte(itemkey),
		Value: []byte(itemvalue),
	})
	if err != nil {
		log.Println("[ERROR] BsPutItem bskey", bskey, "itemkey", itemkey, "err", err)
	}
	log.Println("BsPutItem itemkey", itemkey, "itemvalue", itemvalue, "success")
}
func BsRemoveItem(host, port, bskey, itemkey string) {
	stringbs := StringBigsetService.NewStringBigsetServiceModel2(nil, "/test", host, port)
	err := stringbs.BsRemoveItem(generic.TStringKey(bskey), generic.TItemKey(itemkey))
	if err != nil {
		log.Println("[ERROR] BsRemoveItem bskey", bskey, "itemkey", itemkey, "err", err)
	}
	log.Println("Remove itemkey ", itemkey, "sucesss")
}
func GetTotalCount(host, port, bskey string) {
	stringbs := StringBigsetService.NewStringBigsetServiceModel2(nil, "/test", host, port)
	total, err := stringbs.GetTotalCount(generic.TStringKey(bskey))
	if err != nil {
		log.Println("[ERROR] GetTotalCount bskey", bskey, "err", err)
	}
	log.Println("GetTotalCount bskey ", bskey, "total", total)
}
func GetListKey(host, port string, fromIndex int64, count int32) {
	stringbs := StringBigsetService.NewStringBigsetServiceModel2(nil, "/test", host, port)
	lsKeys, err := stringbs.GetListKey(fromIndex, count)
	if err != nil || lsKeys == nil || len(lsKeys) == 0 {
		log.Println("[ERROR] GetListKey ", "err", err)
	}
	for index, item := range lsKeys {
		log.Println(index, "key", item)
	}
}

func BsGetSlice(host, port string, bskey string, fromPos int32, count int32) {
	stringbs := StringBigsetService.NewStringBigsetServiceModel2(nil, "/test", host, port)
	lsKeys, err := stringbs.BsGetSlice(generic.TStringKey(bskey), fromPos, count)
	if err != nil || lsKeys == nil || len(lsKeys) == 0 {
		log.Println("[ERROR] BsGetSlice ", "err", err)
	}
	for index, item := range lsKeys {
		log.Println(index, "key", string(item.Key), "value", string(item.Value))
	}
}

func BsGetSliceR(host, port string, bskey string, fromPos int32, count int32) {
	stringbs := StringBigsetService.NewStringBigsetServiceModel2(nil, "/test", host, port)
	lsKeys, err := stringbs.BsGetSliceR(generic.TStringKey(bskey), fromPos, count)
	if err != nil || lsKeys == nil || len(lsKeys) == 0 {
		log.Println("[ERROR] BsGetSlice ", "err", err)
	}
	for index, item := range lsKeys {
		log.Println(index, "key", string(item.Key), "value", string(item.Value))
	}
}

func BsRangeQuery(host, port string, bskey string, fromKey string, endKey string) {
	stringbs := StringBigsetService.NewStringBigsetServiceModel2(nil, "/test", host, port)
	lsKeys, err := stringbs.BsRangeQuery(generic.TStringKey(bskey), generic.TItemKey(fromKey), generic.TItemKey(endKey))
	if err != nil || lsKeys == nil || len(lsKeys) == 0 {
		log.Println("[ERROR] BsRangeQuery ", "err", err)
	}
	for index, item := range lsKeys {
		log.Println(index, "key", string(item.Key), "value", string(item.Value))
	}
}
func main() {
	host := flag.String("host", "127.0.0.1", "ip address of service")
	port := flag.String("port", "8883", "port of service")
	action := flag.String("action", "get_item", "action function : BsGetItem,BsPutItem,GetTotalCount,GetTotalStringKey")
	fromIndex := flag.Int64("fromIndex", 0, "use when get list bskey")
	count := flag.Int64("count", 0, "number item want get")
	fromPos := flag.Int64("fromPos", 0, "position in list item")
	bskey := flag.String("bskey", "BSKEY_TEST", "bigset key")
	itemkey := flag.String("itemkey", "ITEM_KEY", "use in getitem")
	itemvalue := flag.String("itemvalue", "ITEM_VALUE", "use when put item")
	startkey := flag.String("startkey", "START_KEY", "begin key")
	endkey := flag.String("endkey", "END_KEY", "end key")
	flag.Parse()
	actionStandar := strings.ToLower(*action)
	switch actionStandar {
	case "bsgetitem":
		BsGetItem(*host, *port, *bskey, *itemkey)
	case "bsputitem":
		BsPutItem(*host, *port, *bskey, *itemkey, *itemvalue)
	case "bsremoveitem":
		BsRemoveItem(*host, *port, *bskey, *itemkey)
	case "gettotalcount":
		GetTotalCount(*host, *port, *bskey)
	case "getlistkey":
		GetListKey(*host, *port, *fromIndex, int32(*count))
	case "bsgetslice":
		BsGetSlice(*host, *port, *bskey, int32(*fromPos), int32(*count))
	case "bsgetslicer":
		BsGetSlice(*host, *port, *bskey, int32(*fromPos), int32(*count))
	case "bsrangequery":
		BsRangeQuery(*host, *port, *bskey, *startkey, *endkey)
	}
}
