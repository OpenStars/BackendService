package KVCounterService

import (
	"context"
	"errors"
	"log"
	"sync"

	etcdconfig "github.com/OpenStars/configetcd"

	"github.com/OpenStars/BackendService/KVCounterService/kvcounter/thrift/gen-go/OpenStars/Counters/KVStepCounter"
	transports "github.com/OpenStars/BackendService/KVCounterService/kvcounter/transportsv2"
)

type KVCounterService struct {
	host   string
	port   string
	sid    string
	schema string
	mu     *sync.RWMutex
}

func (m *KVCounterService) GetValue(genname string) (int64, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	m.mu.Lock()
	client := transports.GetKVCounterCompactClient(m.host, m.port)
	m.mu.Unlock()
	if client == nil || client.Client == nil {
		return -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStepCounter.KVStepCounterServiceClient).GetValue(context.Background(), genname)
	if err != nil {
		transports.ServiceDisconnect(client)
		return -1, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r, nil

}

func (m *KVCounterService) SetValue(genname string, value int64) (int64, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	m.mu.Lock()
	client := transports.GetKVCounterCompactClient(m.host, m.port)
	m.mu.Unlock()
	if client == nil || client.Client == nil {
		return -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStepCounter.KVStepCounterServiceClient).SetValue(context.Background(), genname, value)
	if err != nil {
		transports.ServiceDisconnect(client)
		return -1, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r, nil

}

func (m *KVCounterService) Decrement(genname string, value int64) (int64, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	m.mu.Lock()
	client := transports.GetKVCounterCompactClient(m.host, m.port)
	m.mu.Unlock()
	if client == nil || client.Client == nil {
		return -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStepCounter.KVStepCounterServiceClient).Decrement(context.Background(), genname, value)
	if err != nil {
		transports.ServiceDisconnect(client)
		return -1, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r, nil

}

func (m *KVCounterService) GetMultiStepValue(listKeys []string, step int64) ([]*KVStepCounter.TKVCounterItem, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	m.mu.Lock()
	client := transports.GetKVCounterCompactClient(m.host, m.port)
	m.mu.Unlock()
	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStepCounter.KVStepCounterServiceClient).GetMultiStepValue(context.Background(), listKeys, step)
	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r.ListItems, nil
}

func (m *KVCounterService) GetMultiValue(listKeys []string) ([]*KVStepCounter.TKVCounterItem, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	m.mu.Lock()
	client := transports.GetKVCounterCompactClient(m.host, m.port)
	m.mu.Unlock()
	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStepCounter.KVStepCounterServiceClient).GetMultiValue(context.Background(), listKeys)
	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r.ListItems, nil
}
func (m *KVCounterService) GetMultiCurrentValue(listKeys []string) ([]*KVStepCounter.TKVCounterItem, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	m.mu.Lock()
	client := transports.GetKVCounterCompactClient(m.host, m.port)
	m.mu.Unlock()
	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStepCounter.KVStepCounterServiceClient).GetMultiCurrentValue(context.Background(), listKeys)
	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r.ListItems, nil
}

func (m *KVCounterService) GetCurrentValue(genname string) (int64, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	m.mu.Lock()
	client := transports.GetKVCounterCompactClient(m.host, m.port)
	m.mu.Unlock()

	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStepCounter.KVStepCounterServiceClient).GetCurrentValue(context.Background(), genname)
	if err != nil {
		transports.ServiceDisconnect(client)
		return -1, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r, nil
}

func (m *KVCounterService) GetStepValue(genname string, step int64) (int64, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	m.mu.Lock()
	client := transports.GetKVCounterCompactClient(m.host, m.port)
	m.mu.Unlock()

	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStepCounter.KVStepCounterServiceClient).GetStepValue(context.Background(), genname, step)
	if err != nil {
		transports.ServiceDisconnect(client)
		// client = transports.NewGetKVCounterCompactClient(m.host, m.port)
		return -1, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r, nil
}

func (m *KVCounterService) CreateGenerator(genname string) (int32, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	m.mu.Lock()
	client := transports.GetKVCounterCompactClient(m.host, m.port)
	m.mu.Unlock()

	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStepCounter.KVStepCounterServiceClient).CreateGenerator(context.Background(), genname)
	if err != nil {
		transports.ServiceDisconnect(client)
		return -1, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return r, nil

}

func (m *KVCounterService) RemoveGenerator(genname string) (bool, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	m.mu.Lock()
	client := transports.GetKVCounterCompactClient(m.host, m.port)
	m.mu.Unlock()

	if client == nil || client.Client == nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	_, err := client.Client.(*KVStepCounter.KVStepCounterServiceClient).RemoveGenerator(context.Background(), genname)
	if err != nil {
		transports.ServiceDisconnect(client)
		return false, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer transports.BackToPool(client)
	return true, nil

}
func (m *KVCounterService) WatchChangeEndpoint() {
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
func NewClient(etcdServers []string, sid, defaultHost, defaultPort string) Client {
	ep, _ := etcdconfig.GetEndpoint(sid, "thrift_compact")
	if ep == nil {
		ep = &etcdconfig.Endpoint{
			Host: defaultHost,
			Port: defaultPort,
			SID:  sid,
		}
	}
	log.Println("Init KVCounterService sid", ep.SID, "address", ep.Host+":"+ep.Port)

	client := transports.GetKVCounterCompactClient(ep.Host, ep.Port)
	if client == nil || client.Client == nil {
		log.Println("[ERROR] kvcounter sid", ep.SID, "address", ep.Host+":"+ep.Port, "connection refused")

		return nil
	}
	kvcounter := &KVCounterService{
		host:   ep.Host,
		port:   ep.Port,
		sid:    sid,
		schema: ep.Schema,
		mu:     &sync.RWMutex{},
	}
	go kvcounter.WatchChangeEndpoint()

	// if kvcounter.etcdManager == nil {
	// 	return kvcounter
	// }
	// err = kvcounter.etcdManager.SetDefaultEntpoint(sid, defaultHost, defaultPort)
	// if err != nil {
	// 	log.Println("SetDefaultEndpoint sid", sid, "err", err)
	// 	return nil
	// }
	return kvcounter
}
