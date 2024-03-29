package KVStorageService

import (
	"context"
	"errors"
	"log"
	"sync"

	etcdconfig "github.com/OpenStars/configetcd"

	"github.com/OpenStars/BackendService/KVStorageService/kvstorage/thrift/gen-go/OpenStars/Platform/KVStorage"
	transports "github.com/OpenStars/BackendService/KVStorageService/kvstorage/transportsv2"
)

// update kv
type kvstorageservice struct {
	host   string
	port   string
	sid    string
	schema string
	mu     *sync.RWMutex
	// etcdManager *EndpointsManager.EtcdBackendEndpointManager
}

// func (m *kvstorageservice) notifyEndpointError() {
// 	if m.botClient != nil {
// 		msg := tgbotapi.NewMessage(m.bot_chatID, "Hệ thống kiểm soát phát hiện service kvstorageservice có địa chỉ "+m.host+":"+m.port+" đang không hoạt động")
// 		m.botClient.Send(msg)
// 	}
// }

// func (m *kvstorageservice) Close() {
// 	transports.Close(m.host, m.port)
// }

func (m *kvstorageservice) GetData(key string) (string, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	m.mu.RLock()
	client := transports.GetKVStorageCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return "", errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStorage.KVStorageServiceClient).GetData(context.Background(), key)
	if err != nil {
		transports.ServiceDisconnect(client)
		return "", errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	transports.BackToPool(client)
	if r.ErrorCode != KVStorage.TErrorCode_EGood || r.Data == nil {
		return "", nil
	}
	return r.Data.Value, nil
}

func (m *kvstorageservice) OpenIterate() (int64, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	m.mu.RLock()
	client := transports.GetKVStorageCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStorage.KVStorageServiceClient).OpenIterate(context.Background())
	if err != nil {
		transports.ServiceDisconnect(client)
		return -1, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r, nil
}

func (m *kvstorageservice) CloseIterate(sessionkey int64) error {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	m.mu.RLock()
	client := transports.GetKVStorageCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStorage.KVStorageServiceClient).CloseIterate(context.Background(), sessionkey)
	if err != nil {
		transports.ServiceDisconnect(client)
		return errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	transports.BackToPool(client)
	if r != KVStorage.TErrorCode_EGood {
		return errors.New(r.String())
	}
	return nil
}
func (m *kvstorageservice) NextItem(sessionKey int64) (*KVStorage.KVItem, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	m.mu.RLock()
	client := transports.GetKVStorageCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStorage.KVStorageServiceClient).NextItem(context.Background(), sessionKey)
	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	transports.BackToPool(client)
	if r.ErrorCode != KVStorage.TErrorCode_EGood {
		return nil, errors.New(r.String())
	}
	return r.Data, nil
}

func (m *kvstorageservice) NextListItems(sessionKey, count int64) ([]*KVStorage.KVItem, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	m.mu.RLock()
	client := transports.GetKVStorageCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStorage.KVStorageServiceClient).NexListItems(context.Background(), sessionKey, count)
	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	transports.BackToPool(client)
	if r.ErrorCode != KVStorage.TErrorCode_EGood {
		return nil, errors.New(r.ErrorCode.String())
	}
	return r.Data, nil
}

func (m *kvstorageservice) PutData(key string, value string) (bool, error) {

	m.mu.RLock()
	client := transports.GetKVStorageCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStorage.KVStorageServiceClient).PutData(context.Background(), key, &KVStorage.KVItem{
		Key:   key,
		Value: value,
	})
	if err != nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	transports.BackToPool(client)
	if r != KVStorage.TErrorCode_EGood {
		return false, nil
	}
	return true, nil
}

func (m *kvstorageservice) RemoveData(key string) (bool, error) {

	m.mu.RLock()
	client := transports.GetKVStorageCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStorage.KVStorageServiceClient).RemoveData(context.Background(), key)
	if err != nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	transports.BackToPool(client)
	if r != KVStorage.TErrorCode_EGood {
		return false, nil
	}
	return true, nil
}

func (m *kvstorageservice) GetListData(keys []string) (results []*KVStorage.KVItem, missingkeys []string, err error) {

	m.mu.RLock()
	client := transports.GetKVStorageCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return nil, nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStorage.KVStorageServiceClient).GetListData(context.Background(), keys)
	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, nil, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	transports.BackToPool(client)
	if r.ErrorCode != KVStorage.TErrorCode_EGood {
		return nil, nil, nil
	}

	return r.Data, r.Missingkeys, nil
}

func (m *kvstorageservice) PutMultiData(listData []*KVStorage.KVItem) (err error) {

	m.mu.RLock()
	client := transports.GetKVStorageCompactClient(m.host, m.port)
	m.mu.RUnlock()
	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	_, err = client.Client.(*KVStorage.KVStorageServiceClient).PutMultiData(context.Background(), listData)
	if err != nil {
		transports.ServiceDisconnect(client)
		return errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	transports.BackToPool(client)
	return nil
}

func (m *kvstorageservice) WatchChangeEndpoint() {
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

func NewClient(sid string, host string, port string) Client {
	ep, _ := etcdconfig.GetEndpoint(sid, "thrift_compact")
	if ep == nil {
		ep = &etcdconfig.Endpoint{
			Host: host,
			Port: port,
			SID:  sid,
		}
	}
	log.Println("Init KVStorage Service sid", ep.SID, "address", ep.Host+":"+ep.Port)
	client := transports.GetKVStorageCompactClient(ep.Host, ep.Port)
	if client == nil || client.Client == nil {
		log.Println("[ERROR] kvstorage sid", ep.SID, "address", ep.Host+":"+ep.Port, "connection refused")
		return nil
	}
	kvstorage := &kvstorageservice{
		host:   ep.Host,
		port:   ep.Port,
		sid:    ep.SID,
		schema: ep.Schema,
		mu:     &sync.RWMutex{},
	}
	go kvstorage.WatchChangeEndpoint()
	return kvstorage
}
