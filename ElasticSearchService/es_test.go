package ElasticSearchService

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestES(t *testing.T) {
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
