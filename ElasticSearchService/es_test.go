package ElasticSearchService

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestIndexES(t *testing.T) {
	esClient := NewClient([]string{"http://10.110.1.100:9206"})
	testData := map[string]interface{}{
		"name":   "Lê Hải Sơn",
		"age":    27,
		"avatar": "post.jpg",
		"id":     1,
	}
	testDataBytes, _ := json.Marshal(testData)
	ok, err := esClient.Index("test-data-es", fmt.Sprint(1), string(testDataBytes))
	fmt.Println("index", ok, err)
}

func TestSearchES(t *testing.T) {

	name := "Cao"
	queryString := fmt.Sprintf(`{
		"query": {
		  "multi_match": {
			"query":  "%s",
			"type":   "phrase_prefix",
			"fields": ["name"]
		  } 
		},
		"from": %d,
		"size": %d 
  }`, name, 0, 10)
	esClient := NewClient([]string{"http://10.110.1.100:9206"})
	rawRs, _, err := esClient.SearchRawString("socialnetwork_group", queryString)
	if err != nil {
		log.Fatalln(err)
	}
	for _, raw := range rawRs {
		fmt.Println(string(raw))
	}
}

func TestAggsCount(t *testing.T) {
	esClient := NewClient([]string{"http://10.110.1.100:9206"})
	docCount, err := esClient.AggsNumberRangeByField("index_post_data", "create_time", 1664373500, 1664373718)
	fmt.Println(docCount, err)
}
