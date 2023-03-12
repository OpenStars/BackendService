package QueueService

import (
	"context"
	"errors"
	"log"
	"sync"

	"github.com/OpenStars/BackendService/QueueService/queuedb/thrift/gen-go/Database/QueueDb"
	"github.com/OpenStars/BackendService/QueueService/queuedb/transports"
	"github.com/lehaisonmath6/etcdconfig"
)

type QueueDbClient struct {
	host   string
	port   string
	sid    string
	schema string
	mu     *sync.RWMutex
}

func (m *QueueDbClient) AddItem(queueID string, item *QueueDb.TItem, maxItem int64) (bool, error) {
	m.mu.RLock()
	client := transports.GetInt2ZsetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*QueueDb.QueueDbServiceClient).AddItem(context.Background(), queueID, item, maxItem)
	if err != nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("Int2Zset: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r.Data, nil

}

func (m *QueueDbClient) AddListItem(lsItem []*QueueDb.TItemQueue, maxItem int64) (bool, error) {
	m.mu.RLock()
	client := transports.GetInt2ZsetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	_, err := client.Client.(*QueueDb.QueueDbServiceClient).AddListItems(context.Background(), lsItem, maxItem)
	if err != nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("Int2ZSet: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)

	return true, nil

}

func (m *QueueDbClient) RemoveItem(queueID string, itemID string) (bool, error) {
	m.mu.RLock()
	client := transports.GetInt2ZsetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*QueueDb.QueueDbServiceClient).RemoveItem(context.Background(), queueID, itemID)
	if err != nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("Int2Zset: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r.Data, nil

}
func (m *QueueDbClient) RemoveListItems(lsItems []*QueueDb.TItemQueue) ([]*QueueDb.TItemQueue, error) {
	m.mu.RLock()
	client := transports.GetInt2ZsetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*QueueDb.QueueDbServiceClient).RemoveListItems(context.Background(), lsItems)
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
func (m *QueueDbClient) ListItems(queueID string, offset, limit int32, desc bool) ([]*QueueDb.TItem, int64, error) {
	m.mu.RLock()
	client := transports.GetInt2ZsetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return nil, 0, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*QueueDb.QueueDbServiceClient).ListItems(context.Background(), queueID, offset, limit, desc)
	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, 0, errors.New("Int2Zset: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r.Data, r.Total, nil

}

func (m *QueueDbClient) WatchChangeEndpoint() {
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

func (m *QueueDbClient) GetItem(queueID string, itemID string) (*QueueDb.TItem, error) {
	m.mu.RLock()
	client := transports.GetInt2ZsetCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*QueueDb.QueueDbServiceClient).GetItem(context.Background(), queueID, itemID)
	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("Int2Zset: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	if r.Code == QueueDb.TErrorCode_EGood {
		return r.Data, nil
	}
	return nil, nil

}
func NewClient(etcdServers []string, sid, defaultHost, defaultPort string) *QueueDbClient {

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
	sortedService := &QueueDbClient{
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
