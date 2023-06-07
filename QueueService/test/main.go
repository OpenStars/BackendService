package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/OpenStars/BackendService/QueueService"
	"github.com/OpenStars/BackendService/QueueService/queuedb/thrift/gen-go/Database/QueueDb"
)

var (
	queueDbService = QueueService.NewClient(nil, "/test/sortedset", "0.0.0.0", "8883")
)

func Test1() {
	// setID := -1
	// startTime := time.Now().Unix()
	// // int2ZsetService.AddItemToSet(setID, &SortedSet.TItem{
	// // 	Key:   fmt.Sprint(7),
	// // 	Value: []byte(fmt.Sprint(2)),
	// // 	Score: 115,
	// // })
	// queueDbService.RemoveItem(int64(setID), fmt.Sprint(4))
	// // for i := 0; i < 3; i++ {
	// // 	int2ZsetService.AddItemToSet(setID, &SortedSet.TItem{
	// // 		Key:   fmt.Sprint(i),
	// // 		Value: []byte(fmt.Sprint(i)),
	// // 		Score: int64(i * 2),
	// // 	})
	// // }
	// endTime := time.Now().Unix()
	// fmt.Println("time to put 100000 item", endTime-startTime)

	// lsItems, _, err := queueDbService.ListItems(int64(setID), 0, 100, true)
	// if len(lsItems) == 0 || err != nil {
	// 	log.Fatalln("get list item", err)
	// }
	// fmt.Println("item Size", len(lsItems))
	// for _, item := range lsItems {
	// 	fmt.Println("itemKey", item.Key, "score", item.Score)
	// }
}
func Test0() {
	queueID := "TestLoad0"
	ok, err := queueDbService.AddItem(queueID, &QueueDb.TItem{
		Key:   "lhs1",
		Value: []byte("abc"),
	}, 3)
	ok, err = queueDbService.AddItem(queueID, &QueueDb.TItem{
		Key:   "lhs2",
		Value: []byte("abc"),
	}, 3)
	ok, err = queueDbService.AddItem(queueID, &QueueDb.TItem{
		Key:   "lhs3",
		Value: []byte("abc"),
	}, 3)
	ok, err = queueDbService.AddItem(queueID, &QueueDb.TItem{
		Key:   "lhs4",
		Value: []byte("abc"),
	}, 3)
	fmt.Println(ok, err)

	lsItem, total, err := queueDbService.ListItems(queueID, 0, 1000, false)
	if err != nil || len(lsItem) == 0 {
		log.Fatalln(err)
	}
	for _, item := range lsItem {
		fmt.Println("key", item.Key, "value", string(item.Value), total)
	}
}

func TestLoad() {
	// use 20 go rountine for 10 key
	setIDPrefix := "TestLoad"
	lsqueueID := []string{}
	numQueue := 1000
	numItem := 1000
	for i := 0; i < numQueue; i++ {
		lsqueueID = append(lsqueueID, setIDPrefix+fmt.Sprint(i))
	}

	chanListSetItem := make(chan []*QueueDb.TItemQueue, 1024)
	wg := &sync.WaitGroup{}
	wg.Add(40)
	for i := 0; i < 40; i++ {
		go func(wg *sync.WaitGroup) {
			for listSetItem := range chanListSetItem {
				ok, err := queueDbService.AddListItem(listSetItem, 10)
				if err != nil || !ok {
					log.Println("add list item err", err)
				}
			}
			wg.Done()
		}(wg)
	}
	startTime := time.Now().Unix()
	for i := 0; i < numQueue; i++ {
		lsSetItem := []*QueueDb.TItemQueue{}
		for j := 0; j < numItem; j++ {
			lsSetItem = append(lsSetItem, &QueueDb.TItemQueue{
				QueueID: lsqueueID[j],
				Key:     fmt.Sprint(i) + ":" + fmt.Sprint(j),
				Value:   (fmt.Sprint(i * j)),
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
	// setID := int64(-1)
	// lsItem, _, _ := int2ZsetService.ListItems(setID, 0, 10, true)
	// fmt.Println("len Item", len(lsItem))
	// for _, item := range lsItem {
	// 	fmt.Println("key", item.Key, "value", string(item.Value), "score", item.Score)
	// }
	// fmt.Println(int2ZsetService.GetItem(setID, "lhs6"))
}
func TestRemove() {
	setID := "TestLoad0"
	ok, err := queueDbService.RemoveItem(setID, "997:0")
	fmt.Println(ok, err)
}
func main() {
	// TestRemove()
	// TestLoad()
	// Test0()

	// GetLoad()

}
