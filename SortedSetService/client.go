package SortedSetService

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/OpenStars/BackendService/SortedSetService/sortedsetservice/thrift/gen-go/Database/SortedSet"
	"github.com/OpenStars/BackendService/SortedSetService/sortedsetservice/transports"
	"github.com/lehaisonmath6/etcdconfig"
)

type SortedSetService struct {
	host   string
	port   string
	sid    string
	schema string
	mu     *sync.RWMutex
}

func (m *SortedSetService) AddItemToSet(setID string, item *SortedSet.TItem) (bool, error) {
	m.mu.RLock()
	client := transports.GetSortedSetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*SortedSet.SortedSetServiceClient).AddItemToSet(context.Background(), setID, item)
	if err != nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("SortedSetService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r.Result_, nil

}

func (m *SortedSetService) AddListItem(lsItem []*SortedSet.TItemSet) (bool, error) {
	m.mu.RLock()
	client := transports.GetSortedSetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	_, err := client.Client.(*SortedSet.SortedSetServiceClient).AddListItem(context.Background(), lsItem)
	if err != nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("SortedSetService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return true, nil

}

func (m *SortedSetService) RemoveItem(setID string, itemID string) (bool, error) {
	m.mu.RLock()
	client := transports.GetSortedSetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*SortedSet.SortedSetServiceClient).RemoveItemInSet(context.Background(), setID, itemID)
	if err != nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("SortedSetService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r.Result_, nil

}

func (m *SortedSetService) GetListItem(setID string, offset, limit int32, desc bool) ([]*SortedSet.TItem, error) {
	m.mu.RLock()
	client := transports.GetSortedSetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*SortedSet.SortedSetServiceClient).GetListItem(context.Background(), setID, int16(offset), int16(limit), desc)
	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("SortedSetService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r.LsItems, nil

}

func (m *SortedSetService) WatchChangeEndpoint() {
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
func NewClient(etcdServers []string, sid, defaultHost, defaultPort string) *SortedSetService {

	// ep, _ := etcdconfig.GetEndpoint(sid, "thrift_compact")

	ep := &etcdconfig.Endpoint{
		Host: defaultHost,
		Port: defaultPort,
		SID:  sid,
	}

	log.Println("Init KVCounterService sid", ep.SID, "address", ep.Host+":"+ep.Port)

	client := transports.GetSortedSetCompactClient(ep.Host, ep.Port)
	if client == nil || client.Client == nil {
		log.Println("[ERROR] kvcounter sid", ep.SID, "address", ep.Host+":"+ep.Port, "connection refused")

		return nil
	}
	sortedService := &SortedSetService{
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
