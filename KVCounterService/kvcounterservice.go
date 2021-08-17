package KVCounterService

import (
	"context"
	"errors"

	"github.com/OpenStars/BackendService/KVCounterService/kvcounter/thrift/gen-go/OpenStars/Counters/KVStepCounter"
	transports "github.com/OpenStars/BackendService/KVCounterService/kvcounter/transportsv2"
)

type KVCounterService struct {
	host string
	port string
	sid  string
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
	client := transports.GetKVCounterCompactClient(m.host, m.port)
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
	client := transports.GetKVCounterCompactClient(m.host, m.port)
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
	client := transports.GetKVCounterCompactClient(m.host, m.port)
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
	client := transports.GetKVCounterCompactClient(m.host, m.port)

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

	client := transports.GetKVCounterCompactClient(m.host, m.port)

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

	client := transports.GetKVCounterCompactClient(m.host, m.port)

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

	client := transports.GetKVCounterCompactClient(m.host, m.port)

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

func NewClient(etcdServers []string, sid, defaultHost, defaultPort string) Client {
	client := transports.GetKVCounterCompactClient(defaultHost, defaultPort)
	if client == nil || client.Client == nil {
		return nil
	}
	kvcounter := &KVCounterService{
		host: defaultHost,
		port: defaultPort,
		sid:  sid,
	}

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
