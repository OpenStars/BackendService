package SorlClientService

import (
	"strconv"

	solr "github.com/rtt/Go-Solr"
)

func NewClient(ahost string, aport string, acollection string) Client {

	port, _ := strconv.Atoi(aport)
	s, err := solr.Init(ahost, port, acollection)
	if err != nil {
		return nil
	}
	return &solrclientservice{
		host:       ahost,
		port:       port,
		collection: acollection,
		sconn:      s,
	}
}
