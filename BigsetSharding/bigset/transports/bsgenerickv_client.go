package transports

import (
	"github.com/OpenStars/BackendService/BigsetSharding/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/BackendService/thriftpool"

	"github.com/apache/thrift/lib/go/thrift"
)

var (
	bsGenericMapPool = thriftpool.NewMapPool(1000, 5, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (generic.NewTStringBigSetKVServiceClient(c)) }),
		thriftpool.DefaultClose)

	ibsGenericMapPool = thriftpool.NewMapPool(1000, 5, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (generic.NewTIBSDataServiceClient(c)) }),
		thriftpool.DefaultClose)
)

//GetBsGenericClient client by host:port
func GetBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := bsGenericMapPool.Get(aHost, aPort).Get()
	return client
}
func NewGetBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := bsGenericMapPool.NewGet(aHost, aPort).Get()
	return client
}

func Close(host, port string) {
	bsGenericMapPool.Release(host, port)
}

//GetIBsGenericClient client by host:port
func GetIBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := ibsGenericMapPool.Get(aHost, aPort).Get()
	return client
}

func NewGetIBsGenericClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := ibsGenericMapPool.NewGet(aHost, aPort).Get()
	return client
}
