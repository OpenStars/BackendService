package transports

import (
	"fmt"

	"github.com/OpenStars/BackendService/KVCounterService/kvcounter/thrift/gen-go/OpenStars/Counters/KVStepCounter"
	"github.com/OpenStars/BackendService/thriftpool"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	kvCounterMapPool = thriftpool.NewMapPool(1000, 5, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (KVStepCounter.NewKVStepCounterServiceClient(c)) }),
		thriftpool.DefaultClose)

	kvCounterMapPoolCompact = thriftpool.NewMapPool(1000, 5, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (KVStepCounter.NewKVStepCounterServiceClient(c)) }),
		thriftpool.DefaultClose)
)

func init() {
	fmt.Println("init thrift kvcounter client ")
}

//GetKVCounterBinaryClient client by host:port
func GetKVCounterBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := kvCounterMapPool.Get(aHost, aPort).Get()
	return client
}

//GetKVCounterCompactClient get compact client by host:port
func GetKVCounterCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := kvCounterMapPoolCompact.Get(aHost, aPort).Get()
	return client
}

func Close(host, port string) {
	kvCounterMapPool.Release(host, port)
	kvCounterMapPoolCompact.Release(host, port)
}

// //GetKVCounterBinaryClient client by host:port
// func NewGetKVCounterBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
// 	client, _ := kvCounterMapPool.NewGet(aHost, aPort).Get()
// 	return client
// }

// //GetKVCounterCompactClient get compact client by host:port
// func NewGetKVCounterCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
// 	client, _ := kvCounterMapPoolCompact.NewGet(aHost, aPort).Get()
// 	return client
// }
