package EndpointsManager

import (
	"strings"

	// etcdv3 "github.com/coreos/etcd/clientv3"
	"context"
	"sync"

	etcdv3 "go.etcd.io/etcd/clientv3"
)

// //EnpointManagerIf interface of enpoint manager
// type EnpointManagerIf interface{
// 	GetEndpoint(serviceID string) (host, port string, err error)
// 	SetDefaultEntpoint(serviceID, host, port string) (err error)
// }

//EtcdEnpointManager Endpoint manager for backend service using etcd
type EtcdEndpointManager struct {
	InMemEndpointManager
	EtcdEndpoints []string
	client        *etcdv3.Client

	rootService      string
	EndpointsMap     sync.Map //value is an array of Endpoint : []*Endpoint
	EndpointRotating sync.Map // map from serviceID to int
}

func (o *EtcdEndpointManager) getFromEtcd(serviceID string) (host, port string, err error) {
	if o.client != nil {
		// try to get from etcd
		// resp, gerr := o.client.Get(context.Background(), serviceID)
		resp, gerr := o.client.Get(context.Background(), serviceID)
		if gerr == nil {
			for _, kv := range resp.Kvs {
				if string(kv.Key) == serviceID {
					Eps := o.parseServiceFromString(serviceID, string(kv.Value))
					if len(Eps) > 0 {
						host, port, err = Eps[0].Host, Eps[0].Port, nil
						return
					}
				}
			}

		}
		ch := o.client.Watch(context.Background(), serviceID, nil...)
		go o.MonitorChan(ch)
	}
	return "", "", nil
}

//GetEndpoint (serviceID string) (host, port string, err error)
func (o *EtcdEndpointManager) GetEndpoint(serviceID string) (host, port string, err error) {
	if o.client == nil {
		o.Start()
	}
	/*
		Get in Endpoint map first
		If it does not exist, get from etcd, and monitor it.
	*/
	endpoints, ok := o.EndpointsMap.Load(serviceID)
	rotated, _ := o.EndpointRotating.LoadOrStore(serviceID, 0)
	o.EndpointRotating.Store(serviceID, rotated.(int)+1)
	if ok {
		arr := endpoints.([]*Endpoint)
		if len(arr) > 0 {
			iPos := rotated.(int) % len(arr)
			host, port = arr[iPos].Host, arr[iPos].Port
			err = nil
			return
		}
		return o.InMemEndpointManager.GetEndpoint(serviceID)

	}

	host, port, err = o.getFromEtcd(serviceID)
	if host != "" {
		return
	}
	// if o.client != nil {
	// 	// try to get from etcd
	// 	// resp, gerr := o.client.Get(context.Background(), serviceID)
	// 	resp, gerr := o.client.Get(context.Background(), serviceID)
	// 	if gerr == nil {
	// 		for _, kv := range resp.Kvs {
	// 			if string(kv.Key) == serviceID {
	// 				Eps := o.parseServiceFromString(serviceID, string(kv.Value))
	// 				if len(Eps) > 0 {
	// 					host, port, err = Eps[0].Host, Eps[0].Port, nil
	// 					return
	// 				}
	// 			}
	// 		}

	// 	}
	// 	ch := o.client.Watch(context.Background(), serviceID, nil...)
	// 	go o.MonitorChan(ch)

	// 	if gerr == nil {
	// 	} else {
	// 		return o.InMemEndpointManager.GetEndpoint(serviceID)
	// 	}
	// }

	//Get from default endpoint
	return o.InMemEndpointManager.GetEndpoint(serviceID)

}

//Get all Endpoint of a serviceID
func (o *EtcdEndpointManager) GetAllEndpoint(serviceID string) ([]*Endpoint, error) {
	if o.client == nil {
		o.Start()
	}
	/*
		Get in Endpoint map first
		If it does not exist, get from etcd, and monitor it.
	*/
	endpoints, ok := o.EndpointsMap.Load(serviceID)
	if ok {
		eps := endpoints.([]*Endpoint)
		if len(eps) > 0 {
			return eps, nil
		}

	} else {
		o.getFromEtcd(serviceID)
		endpoints, ok = o.EndpointsMap.Load(serviceID)
		eps := endpoints.([]*Endpoint)
		if len(eps) > 0 {
			return eps, nil
		}
	}
	return o.InMemEndpointManager.GetAllEndpoint(serviceID)
}

//SetDefaultEntpoint Set default endpoint manager
func (o *EtcdEndpointManager) SetDefaultEntpoint(serviceID, host, port string) (err error) {
	o.InMemEndpointManager.SetDefaultEntpoint(serviceID, host, port)

	if o.client != nil {
		//Already connected to etcdserver
		o.GetEndpoint(serviceID)
	}

	return
}

//NewEtcdEndpointManager Create endpoint manager
func NewEtcdEndpointManager(etcdConfigHostports []string) *EtcdEndpointManager {

	o := &EtcdEndpointManager{
		InMemEndpointManager: InMemEndpointManager{
			defaultEndpoints: make(map[string]*Endpoint),
		},
		EtcdEndpoints: etcdConfigHostports,
		client:        nil,
	}

	return o
}

func (o *EtcdEndpointManager) Start() bool {
	if o.client != nil {
		return false
	}

	if len(o.EtcdEndpoints) == 0 {
		return false
	}

	cfg := etcdv3.Config{
		Endpoints: o.EtcdEndpoints,
	}
	aClient, err := etcdv3.New(cfg)
	if err != nil {
		return false
	}
	o.client = aClient

	opts := []etcdv3.OpOption{etcdv3.WithPrefix()}
	if len(o.rootService) > 0 {

		rootWatchan := aClient.Watch(context.Background(), o.rootService, opts...)
		go o.MonitorChan(rootWatchan)
	}

	for serviceID := range o.defaultEndpoints {
		serviceChan := aClient.Watch(context.Background(), serviceID, nil...)
		go o.MonitorChan(serviceChan)
	}

	return true
}

func (o *EtcdEndpointManager) parseServiceFromString(serviceID, val string) []*Endpoint {
	HostPorts := strings.Split(val, ",")
	var Eps []*Endpoint
	for _, HostPort := range HostPorts {
		hp := strings.Split(HostPort, ":")
		if len(hp) == 2 {
			aHost := strings.Trim(hp[0], " ")
			aPort := strings.Trim(hp[1], " ")
			Eps = append(Eps, &Endpoint{aHost, aPort, 0, ""})
		}
	}
	o.EndpointsMap.Store(serviceID, Eps)
	return Eps
}

//MonitorChan monitor an etcd watcher channel
func (o *EtcdEndpointManager) MonitorChan(wchan etcdv3.WatchChan) {
	for wresp := range wchan {
		for _, ev := range wresp.Events {
			//fmt.Printf("Watch V3 .... %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			if ev.Type == etcdv3.EventTypePut {
				val := string(ev.Kv.Value)
				serviceID := string(ev.Kv.Key)
				o.parseServiceFromString(serviceID, val)
			}
		}
	}
}
