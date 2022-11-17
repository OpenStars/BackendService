package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	sortedsetservice "github.com/OpenStars/BackendService/SortedSetService"
	"github.com/OpenStars/BackendService/SortedSetService/sortedsetservice/thrift/gen-go/Database/SortedSet"
)

var (
	sortedSetService = sortedsetservice.NewClient(nil, "/test/sortedset", "0.0.0.0", "8883")
)

func Test1() {
	setID := "TestMax3"
	startTime := time.Now().Unix()
	// sortedSetService.AddItemToSet(setID, &SortedSet.TItem{
	// 	Key:   fmt.Sprint(7),
	// 	Value: []byte(fmt.Sprint(2)),
	// 	Score: 115,
	// })
	sortedSetService.RemoveItem(setID, fmt.Sprint(4))
	// for i := 0; i < 3; i++ {
	// 	sortedSetService.AddItemToSet(setID, &SortedSet.TItem{
	// 		Key:   fmt.Sprint(i),
	// 		Value: []byte(fmt.Sprint(i)),
	// 		Score: int64(i * 2),
	// 	})
	// }
	endTime := time.Now().Unix()
	fmt.Println("time to put 100000 item", endTime-startTime)

	lsItems, err := sortedSetService.GetListItem(setID, 0, 100, true)
	if len(lsItems) == 0 || err != nil {
		log.Fatalln("get list item", err)
	}
	fmt.Println("item Size", len(lsItems))
	for _, item := range lsItems {
		fmt.Println("itemKey", item.Key, "score", item.Score)
	}
}
func Test0() {
	setID := "Test01"
	ok, err := sortedSetService.AddItemToSet(setID, &SortedSet.TItem{
		Key:   "lhs",
		Value: []byte("abc"),
		Score: 250,
	})
	fmt.Println(ok, err)

	lsItem, err := sortedSetService.GetListItem(setID, 0, 1000, true)
	if err != nil || len(lsItem) == 0 {
		log.Fatalln(err)
	}
	for _, item := range lsItem {
		fmt.Println("key", item.Key, "score", item.Score)
	}
}

func TestLoad() {
	// use 20 go rountine for 10 key
	setIDPrefix := "TestLoad"
	lsSetID := []string{}

	for i := 0; i < 100; i++ {
		lsSetID = append(lsSetID, setIDPrefix+fmt.Sprint(i))
	}

	chanListSetItem := make(chan []*SortedSet.TItemSet, 1024)
	wg := &sync.WaitGroup{}
	wg.Add(40)
	for i := 0; i < 40; i++ {
		go func(wg *sync.WaitGroup) {
			for listSetItem := range chanListSetItem {
				ok, err := sortedSetService.AddListItem(listSetItem)
				if err != nil || !ok {
					log.Println("add list item err", err)
				}
			}
			wg.Done()
		}(wg)
	}
	startTime := time.Now().Unix()
	for i := 0; i < 100; i++ {
		lsSetItem := []*SortedSet.TItemSet{}
		for j := 0; j < 100; j++ {
			lsSetItem = append(lsSetItem, &SortedSet.TItemSet{
				SetID: lsSetID[j],
				Key:   fmt.Sprint("i"),
				Value: []byte(fmt.Sprint(i)),
				Score: rand.Int63(),
			})
		}
		chanListSetItem <- lsSetItem
	}
	close(chanListSetItem)
	wg.Wait()
	endTime := time.Now().Unix()
	fmt.Println("total test time", endTime-startTime)
}

func GetLoad() {
	setID := "TestLoad9"
	lsItem, _ := sortedSetService.GetListItem(setID, 0, 10, true)
	fmt.Println("len Item", len(lsItem))
	for _, item := range lsItem {
		fmt.Println("key", item.Key, "value", string(item.Value), "score", item.Score)
	}
}
func main() {
	TestLoad()
	Test0()
	GetLoad()
}
