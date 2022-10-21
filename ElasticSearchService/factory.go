package ElasticSearchService

import (
	"log"

	"github.com/elastic/go-elasticsearch/v7"
)

func NewClient(hosts []string) Client {
	log.Println("Init ElasticSearchClient address", hosts)
	cfg := elasticsearch.Config{
		Addresses: hosts,
	}
	url := hosts[0]
	if hosts[0][len(hosts[0])-1] == '/' {
		url = hosts[0][:len(hosts[0])-1]
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Println("[ERROR] NewESClient err", err)
		return nil
	}
	return &client{
		esclient: es,
		url:      url,
	}

}
