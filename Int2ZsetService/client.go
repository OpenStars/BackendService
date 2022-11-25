package Int2ZsetService

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/OpenStars/BackendService/Int2ZsetService/int2zset/thrift/gen-go/Database/Int2Zset"
	"github.com/OpenStars/BackendService/Int2ZsetService/int2zset/transports"
	"github.com/lehaisonmath6/etcdconfig"
)

type int2zset struct {
	host   string
	port   string
	sid    string
	schema string
	mu     *sync.RWMutex
}

func (m *int2zset) AddItem(setID int64, item *Int2Zset.TItem, maxItem int64) (bool, error) {
	m.mu.RLock()
	client := transports.GetInt2ZsetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*Int2Zset.Int2ZsetServiceClient).AddItem(context.Background(), setID, item, maxItem)
	if err != nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("Int2Zset: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r.Data, nil

}

func (m *int2zset) AddListItem(lsItem []*Int2Zset.TItemSet, maxItem int64) (bool, error) {
	m.mu.RLock()
	client := transports.GetInt2ZsetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	_, err := client.Client.(*Int2Zset.Int2ZsetServiceClient).AddListItems(context.Background(), lsItem, maxItem)
	if err != nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("Int2ZSet: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)

	return true, nil

}

func (m *int2zset) RemoveItem(setID int64, itemID string) (bool, error) {
	m.mu.RLock()
	client := transports.GetInt2ZsetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*Int2Zset.Int2ZsetServiceClient).RemoveItem(context.Background(), setID, itemID)
	if err != nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("Int2Zset: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r.Data, nil

}
func (m *int2zset) RemoveListItems(lsItems []*Int2Zset.TItemSet) ([]*Int2Zset.TItemSet, error) {
	m.mu.RLock()
	client := transports.GetInt2ZsetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*Int2Zset.Int2ZsetServiceClient).RemoveListItems(context.Background(), lsItems)
	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("Int2Zset: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	if len(r.Data) == 0 {
		return nil, nil
	}
	return r.Data, nil

}
func (m *int2zset) ListItems(setID int64, offset, limit int32, desc bool) ([]*Int2Zset.TItem, int64, error) {
	m.mu.RLock()
	client := transports.GetInt2ZsetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return nil, 0, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*Int2Zset.Int2ZsetServiceClient).ListItems(context.Background(), setID, offset, limit, desc)
	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, 0, errors.New("Int2Zset: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r.Data, r.Total, nil

}

func (m *int2zset) WatchChangeEndpoint() {
	epChan := make(chan *etcdconfig.Endpoint)
	go etcdconfig.WatchChangeService(m.sid, epChan)
	for ep := range epChan {
		log.Println("[EVENT CHANGE ENDPOINT] sid", m.sid, "to address", ep.Host+":"+ep.Port)
		m.mu.Lock()
		m.host = ep.Host
		m.port = ep.Port
		m.mu.Unlock()
	}
}
func NewClient(etcdServers []string, sid, defaultHost, defaultPort string) *int2zset {

	// ep, _ := etcdconfig.GetEndpoint(sid, "thrift_compact")

	ep := &etcdconfig.Endpoint{
		Host: defaultHost,
		Port: defaultPort,
		SID:  sid,
	}

	log.Println("Init Zset sid", ep.SID, "address", ep.Host+":"+ep.Port)

	client := transports.GetInt2ZsetCompactClient(ep.Host, ep.Port)
	if client == nil || client.Client == nil {
		log.Println("[ERROR] zset sid", ep.SID, "address", ep.Host+":"+ep.Port, "connection refused")

		return nil
	}
	sortedService := &int2zset{
		host:   ep.Host,
		port:   ep.Port,
		sid:    sid,
		schema: ep.Schema,
		mu:     &sync.RWMutex{},
	}
	if etcdServers != nil {
		go sortedService.WatchChangeEndpoint()
	}

	// if kvcounter.etcdManager == nil {
	// 	return kvcounter
	// }
	// err = kvcounter.etcdManager.SetDefaultEntpoint(sid, defaultHost, defaultPort)
	// if err != nil {
	// 	log.Println("SetDefaultEndpoint sid", sid, "err", err)
	// 	return nil
	// }
	return sortedService
}
