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

	name := "Sơn"
	queryString := fmt.Sprintf(`{
		"query": {
		  "multi_match": {
			"query":  "%s",
			"type":   "phrase_prefix",
			"fields": ["display_name"]
		  } 
		},
		"from": %d,
		"size": %d 
  }`, name, 0, 10)
	// 	queryString := fmt.Sprintf(`{
	// 		"query": {
	// 		  "multi_match": {
	// 			"query":  "%s",
	// 			"type":   "phrase_prefix",
	// 			"fields": ["name"]
	// 		  }
	// 		},
	// 		"from": %d,
	// 		"size": %d
	//   }`, name, 0, 10)
	esClient := NewClient([]string{"http://10.110.1.100:9206"})
	rawRs, _, err := esClient.SearchRawString("index_profile_info_by_appid_trustkeys", queryString)
	if err != nil {
		log.Fatalln(err)
	}
	for _, raw := range rawRs {
		fmt.Println(string(raw))
	}
}

func TestIndex(t *testing.T) {
	esClient := NewClient([]string{"http://10.110.1.100:9206"})
	docJson := `{"id":100,"name":"USD Coin","symbol":"USDC","decimal":6,"precision":6,"logo":"https://basescan.org/token/images/centre-usdc_28.png","category":"token","price_usd":{"price":1,"volume_24h":100,"percent_change_1h":0,"percent_change_24h":0,"percent_change_7d":0,"percent_change_30d":0,"percent_change_60d":0,"percent_change_90d":0,"market_cap":0,"fully_diluted_market_cap":0,"market_cap_by_total_supply":0,"dominance":0,"turnover":0,"ytd_price_change_percentage":0},"blockchain":{"id":23,"name":"Base Mainnet","symbol":"ETH(BASE)","decimal":18,"adapter":"ETH","chain_id":8453,"network_type":"mainnet","crypto_id":87,"symbol_crypto":"ETH(BASE)","logo":"https://s2.coinmarketcap.com/static/img/coins/64x64/27716.png"},"token":{"id":100,"name":"USD Coin","symbol":"USDC","decimal":6,"blockchain":{"id":23,"name":"Base Mainnet","symbol":"ETH(BASE)","decimal":18,"adapter":"ETH","chain_id":8453,"network_type":"mainnet","crypto_id":87,"symbol_crypto":"ETH(BASE)","logo":"https://s2.coinmarketcap.com/static/img/coins/64x64/27716.png"},"address":"0x833589fcd6edb6e08f4c7c32d4f71b54bda02913","logo":"https://basescan.org/token/images/centre-usdc_28.png","standard":""},"crawl_cryptocurrency_id":0,"create_time":1711972461,"origin_crypto_id":0,"explore_address_url":"https://basescan.com/address/","explores":[],"description":"","circulating_supply":100,"total_supply":0,"max_supply":10000,"org":null,"map_language_description":{},"white_paper":"https://circle.com"}`
	esClient.Index("cryptocurrency-info", "100", docJson)
}
func TestAggsCount(t *testing.T) {
	esClient := NewClient([]string{"http://10.110.1.100:9206"})
	docCount, err := esClient.AggsNumberRangeByField("index_post_data", "create_time", 1664373500, 1664373718)
	fmt.Println(docCount, err)
}
