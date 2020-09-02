package EndpointsManager

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"context"
	"sync"

	etcdv3 "go.etcd.io/etcd/clientv3"
)

//EtcdEnpointManager Endpoint manager for backend service using etcd
type EtcdBackendEndpointManager struct {
	InMemEndpointManager
	EtcdEndpoints []string
	client        *etcdv3.Client

	rootService      string
	EndpointsMap     sync.Map //value is an array of Endpoint : []*Endpoint
	EndpointRotating sync.Map // map from serviceID to int
}

//parseEndpointFromPath
//Input: a full path /my/path/thrift:10.0.0.10:5560
//output : errr, endpoint , a service path (/my/path)
func (e *EtcdBackendEndpointManager) parseEndpointFromPath(endPointPath string) (error, *Endpoint, string) {
	var ep Endpoint
	baseNode := strings.Split(endPointPath, "/")
	fmt.Println(endPointPath, " base Node", baseNode)
	if len(baseNode) == 0 {
		return errors.New("Parse endpoint error " + endPointPath), nil, ""
	}
	nodeName := baseNode[len(baseNode)-1] //Token cuối cùng, format schema:host:port
	fmt.Println("node name: ", nodeName)
	token := strings.Split(nodeName, ":")

	if len(token) != 3 {
		return errors.New("Parse endpoint error " + nodeName), nil, ""
	}

	port := token[2]

	ep.Type = StringToTType(token[0])
	ep.Host = token[1]
	ep.Port = port
	serviceID := strings.Join(baseNode[:len(baseNode)-1], "/")
	ep.ServiceID = serviceID
	fmt.Println("parsed ", ep, serviceID)
	return nil, &ep, serviceID
}

func (o *EtcdBackendEndpointManager) getFromEtcd(serviceID string) (host, port string, err error) {
	if o.client != nil {
		// try to get from etcd
		// resp, gerr := o.client.Get(context.Background(), serviceID)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		ch := o.client.Watch(ctx, serviceID, nil...)

		go o.MonitorChan(ch)

		opts := []etcdv3.OpOption{etcdv3.WithPrefix()}
		ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
		resp, gerr := o.client.Get(ctx, serviceID, opts...)
		cancel()
		if gerr == nil {
			var Eps []*Endpoint

			for _, kv := range resp.Kvs {
				// fmt.Println("got ", string(kv.Key), serviceID)
				if strings.HasPrefix(string(kv.Key), serviceID) {
					// fmt.Println("got ", string(kv.Key), serviceID)
					_, anEp, parsedServiceID := o.parseEndpointFromPath(string(kv.Key))
					if anEp != nil && parsedServiceID == serviceID {
						Eps = append(Eps, anEp)
						host = anEp.Host
						port = anEp.Port
						err = nil
					}
				}
			}

			o.EndpointsMap.Store(serviceID, Eps)
			if len(Eps) > 0 {
				return
			}

		}

	}

	return "", "", nil
}

//GetEndpoint (serviceID string) (host, port string, err error)
func (o *EtcdBackendEndpointManager) GetEndpoint(serviceID string) (host, port string, err error) {

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
		// fmt.Println("get from store")
		arr := endpoints.([]*Endpoint)
		if len(arr) > 0 {
			iPos := rotated.(int) % len(arr)
			host, port = arr[iPos].Host, arr[iPos].Port
			err = nil
			return
		}
		return o.InMemEndpointManager.GetEndpoint(serviceID)

	}

	//Get from default endpoint
	return o.InMemEndpointManager.GetEndpoint(serviceID)

}

//Get all Endpoint of a serviceID
func (o *EtcdBackendEndpointManager) GetAllEndpoint(serviceID string) ([]*Endpoint, error) {
	// Do nothing
	// if o.client == nil {
	// 	o.Start()
	// }
	/*
		Get in Endpoint map first
		If it does not exist, get from etcd, and monitor it.
	// */
	// endpoints, ok := o.EndpointsMap.Load(serviceID)
	// if ok {
	// 	eps := endpoints.([]*Endpoint)
	// 	if len(eps) > 0 {
	// 		return eps, nil
	// 	}
	// }

	// } else {
	// 	o.getFromEtcd(serviceID)
	// 	endpoints, ok = o.EndpointsMap.Load(serviceID)
	// 	eps := endpoints.([]*Endpoint)
	// 	if len(eps) > 0 {
	// 		return eps, nil
	// 	}
	// }
	// return o.InMemEndpointManager.GetAllEndpoint(serviceID)
	return nil, nil
}

func (o *EtcdBackendEndpointManager) GetAllEndPoint(serviceID string) ([]*Endpoint, error) {
	if o.client == nil {
		o.Start()
	}
	/*
		Get in Endpoint map first
		If it does not exist, get from etcd, and monitor it.
	// */
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
func (o *EtcdBackendEndpointManager) SetDefaultEntpoint(serviceID, host, port string) (err error) {
	log.Println("EtcdBackendEndpointManager SetDefaultEndpoint ", serviceID, host, port)
	o.InMemEndpointManager.SetDefaultEntpoint(serviceID, host, port)

	if o.client != nil {
		o.getFromEtcd(serviceID)
	}
	// if o.client != nil {
	// 	//Already connected to etcdserver
	// 	o.GetEndpoint(serviceID)
	// }

	return
}

//NewEtcdBackendEndpointManager Create endpoint manager
func NewEtcdBackendEndpointManager(etcdConfigHostports []string) *EtcdBackendEndpointManager {

	o := &EtcdBackendEndpointManager{
		InMemEndpointManager: InMemEndpointManager{
			defaultEndpoints: make(map[string]*Endpoint),
		},
		EtcdEndpoints: etcdConfigHostports,
		client:        nil,
	}

	return o
}

func (o *EtcdBackendEndpointManager) Start() bool {
	fmt.Println("Starting Backend Endpoint manager etcd  ", o.EtcdEndpoints)
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
		fmt.Println("Start Watching ", serviceID)
		serviceChan := aClient.Watch(context.Background(), serviceID, opts...)
		go o.MonitorChan(serviceChan)
		go o.getFromEtcd(serviceID)
	}

	return true
}

// func (o *EtcdBackendEndpointManager) parseServiceFromString(serviceID, val string) []*Endpoint {
// 	HostPorts := strings.Split(val, ",")
// 	var Eps []*Endpoint
// 	for _, HostPort := range HostPorts {
// 		hp := strings.Split(HostPort, ":")
// 		if len(hp) == 2 {
// 			aHost := strings.Trim(hp[0], " ")
// 			aPort := strings.Trim(hp[1], " ")
// 			Eps = append(Eps, &Endpoint{aHost, aPort, 0, ""})
// 		}
// 	}
// 	o.EndpointsMap.Store(serviceID, Eps)
// 	return Eps
// }

func (o *EtcdBackendEndpointManager) addEndpoint(ep *Endpoint, serviceID string) {
	endpoints, _ := o.EndpointsMap.Load(serviceID)
	var Eps []*Endpoint
	if endpoints != nil {
		Eps = endpoints.([]*Endpoint)
	}
	// check already exist
	for _, aEp := range Eps {
		if aEp.Host == ep.Host && aEp.Port == ep.Port {
			return
		}
	}
	Eps = append(Eps, ep)
	o.EndpointsMap.Store(serviceID, Eps)
}

func (o *EtcdBackendEndpointManager) removeEndpoint(ep *Endpoint, serviceID string) {
	endpoints, _ := o.EndpointsMap.Load(serviceID)
	var Eps []*Endpoint
	if endpoints != nil {
		Eps = endpoints.([]*Endpoint)
	}

	var newEps []*Endpoint
	// check already exist
	for _, aEp := range Eps {
		if aEp.Host == ep.Host && aEp.Port == ep.Port {
			continue
		}
		newEps = append(newEps, aEp)
	}

	o.EndpointsMap.Store(serviceID, newEps)
}

//MonitorChan monitor an etcd watcher channel
func (o *EtcdBackendEndpointManager) MonitorChan(wchan etcdv3.WatchChan) {
	for wresp := range wchan {
		for _, ev := range wresp.Events {
			//fmt.Printf("Watch V3 .... %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
			if ev.Type == etcdv3.EventTypePut {
				// val := string(ev.Kv.Value)
				serviceFullPath := string(ev.Kv.Key)
				// o.parseServiceFromString(serviceID, val)
				err, ep, serviceID := o.parseEndpointFromPath(serviceFullPath)
				if err == nil {
					o.addEndpoint(ep, serviceID)
				}

			} else if ev.Type == etcdv3.EventTypeDelete {
				serviceFullPath := string(ev.Kv.Key)
				// o.parseServiceFromString(serviceID, val)
				err, ep, serviceID := o.parseEndpointFromPath(serviceFullPath)
				fmt.Println("Delete ", serviceFullPath, " ", string(ev.Kv.Key))
				if err == nil {
					o.removeEndpoint(ep, serviceID)
				}

			}
		}
	}
}
