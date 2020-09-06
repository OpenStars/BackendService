package SorlClientService

import (
	solr "github.com/rtt/Go-Solr"
)

func NewSolrClient(ahost string, aport int, acollection string) SolrClientServiceIf {

	s, err := solr.Init(ahost, aport, acollection)
	if err != nil {
		return nil
	}
	return &solrclientservice{
		host:       ahost,
		port:       aport,
		collection: acollection,
		sconn:      s,
	}
}
