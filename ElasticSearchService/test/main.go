package main

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/OpenStars/BackendService/ElasticSearchService"
)

type ShopItem struct {
	ID           int64   `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Price        int64   `json:"price,omitempty"`
	Discount     float64 `json:"discount,omitempty"`
	CategoryID   int64   `json:"category_id,omitempty"`
	CategoryName string  `json:"category_name,omitempty"`
	Description  string  `json:"description,omitempty"`
	Timestamp    int64   `json:"timestamp,omitempty"`
}

func IndexToES(esclient ElasticSearchService.Client) {
	sitem1 := ShopItem{
		ID:           1,
		Name:         "Áo sơ mi trắng nam",
		Price:        300000,
		Discount:     0,
		CategoryID:   1,
		CategoryName: "Áo nam",
		Description:  "Áo sơ mi trắng nam phù hợp cho giới trẻ trong mùa hè này, thích hợp đi chơi dự liên hoan lớp",
		Timestamp:    1598342499,
	}
	databytes, _ := json.Marshal(sitem1)
	id, err := esclient.Update("eshop", strconv.Itoa(1), string(databytes))
	log.Println("Index success id", id)
	if err != nil {
		log.Println("[ERROR] Index docID", sitem1.ID, "err", err)
	}
	sitem2 := ShopItem{
		ID:           2,
		Name:         "Áo phông google nam",
		Price:        150000,
		Discount:     0,
		CategoryID:   1,
		CategoryName: "Áo nam",
		Description:  "Áo google cho nam phù hợp với dân IT và là fan hâm mộ của google",
		Timestamp:    1598256099,
	}

	databytes, _ = json.Marshal(sitem2)
	id, err = esclient.Index("eshop", strconv.Itoa(2), string(databytes))
	log.Println("Index success id", id)
	if err != nil {
		log.Println("[ERROR] Index docID", sitem2.ID, "err", err)
	}

	sitem3 := ShopItem{
		ID:           3,
		Name:         "Áo phông google nữ",
		Price:        150000,
		Discount:     0,
		CategoryID:   1,
		CategoryName: "Áo nữ",
		Description:  "Áo google cho nữ phù hợp với dân IT và là fan hâm mộ của google",
		Timestamp:    1598169699,
	}

	databytes, _ = json.Marshal(sitem3)
	id, err = esclient.Index("eshop", strconv.Itoa(3), string(databytes))
	log.Println("Index success id", id)
	if err != nil {
		log.Println("[ERROR] Index docID", sitem3.ID, "err", err)
	}
}

func GetES(esclient ElasticSearchService.Client) {

}

func main() {

	// urlPath := "http://10.110.1.21:9092/"
	// log.Println(urlPath[])
	// urlPath := "http://10.110.1.21:9092/"
	// rawUrl, _ := url.Parse(urlPath)
	// log.Println(rawUrl.Path)
	esclient := ElasticSearchService.NewClient([]string{"http://10.110.1.100:9206/"})

	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"displayName": "Sơn",
			},
		},
	}
	rs, err := esclient.Search("trustkeysprofile", query)

	if err != nil {
		log.Fatalln(err)
	}
	doc, err := ElasticSearchService.ParseResultToDocuments(rs)
	// IndexToES(esclient)
	// result, err := esclient.Get("eshop", "qUl2LXQBTjAo_D0k19Q7")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(doc)
}
