package ElasticSearchService

import (
	"log"

	"github.com/elastic/go-elasticsearch"
)

func NewClient(hosts []string) Client {
	cfg := elasticsearch.Config{
		Addresses: hosts,
	}
	url := hosts[0]
	if hosts[0][len(hosts[0])-1] == '/' {
		url = hosts[0][:len(hosts[0])-2]
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
