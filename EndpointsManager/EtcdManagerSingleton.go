package EndpointsManager

import (
	"fmt"
	"log"
	"sync"
	"time"

	etcdv3 "go.etcd.io/etcd/v3"
)

var etcdManagerSingleton *EtcdBackendEndpointManager
var once sync.Once

func GetEtcdBackendEndpointManagerSingleton(etcdEndpoints []string) *EtcdBackendEndpointManager {
	once.Do(func() {
		etcdManagerSingleton = NewEtcdBackendEndpointManager(etcdEndpoints)
		fmt.Println("Starting Backend Endpoint manager etcd  ", etcdEndpoints)

		if len(etcdEndpoints) == 0 {
			etcdManagerSingleton = nil
			return
		}

		cfg := etcdv3.Config{
			Endpoints:   etcdEndpoints,
			DialTimeout: 3 * time.Second,
		}
		aClient, err := etcdv3.New(cfg)
		if err != nil {
			log.Println("Etcd client err", err)
			etcdManagerSingleton = nil
		}

		if etcdManagerSingleton != nil {
			etcdManagerSingleton.client = aClient
		}
	})
	return etcdManagerSingleton
}
