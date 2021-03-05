package transports

import (
	"fmt"

	"github.com/OpenStars/BackendService/KVStorageService/kvstorage/thrift/gen-go/OpenStars/Platform/KVStorage"
	"github.com/OpenStars/BackendService/thriftpool"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	kvstorageMapPool = thriftpool.NewMapPool(1000, 5, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (KVStorage.NewKVStorageServiceClient(c)) }),
		thriftpool.DefaultClose)

	kvstorageMapPoolCompact = thriftpool.NewMapPool(1000, 5, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (KVStorage.NewKVStorageServiceClient(c)) }),
		thriftpool.DefaultClose)
)

func init() {
	fmt.Println("init thrift kvcounter client ")
}

//GetKVCounterBinaryClient client by host:port
func GetKVStorageBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := kvstorageMapPool.Get(aHost, aPort).Get()
	return client
}

//GetKVCounterCompactClient get compact client by host:port
func GetKVStorageCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := kvstorageMapPoolCompact.Get(aHost, aPort).Get()
	return client
}

func Close(host, port string) {
	kvstorageMapPool.Release(host, port)
	kvstorageMapPoolCompact.Release(host, port)
}
