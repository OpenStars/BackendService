package transports

import (
	"github.com/OpenStars/BackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/listi64"
	"github.com/OpenStars/BackendService/thriftpool"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	stringBsListI64MapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (listi64.NewTStringBSListI64ServiceClient(c)) }),
		thriftpool.DefaultClose)

	i64BsListI64MapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (listi64.NewTIBSListI64ServiceClient(c)) }),
		thriftpool.DefaultClose)
)

//GetStringBsListI64Client client by host:port
func GetStringBsListI64Client(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := stringBsListI64MapPool.Get(aHost, aPort).Get()
	return client
}

//GetI64BsListI64Client client by host:port
func GetI64BsListI64Client(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := i64BsListI64MapPool.Get(aHost, aPort).Get()
	return client
}
