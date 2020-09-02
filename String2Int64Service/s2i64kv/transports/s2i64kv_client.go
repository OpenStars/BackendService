package transports

import (
	"fmt"

	"github.com/OpenStars/BackendService/String2Int64Service/s2i64kv/thrift/gen-go/OpenStars/Common/S2I64KV"
	"github.com/OpenStars/BackendService/thriftpool"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	s2i64BinMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (S2I64KV.NewTString2I64KVServiceClient(c)) }),
		thriftpool.DefaultClose)

	s2I64CompactMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (S2I64KV.NewTString2I64KVServiceClient(c)) }),
		thriftpool.DefaultClose)
)

func init() {
	fmt.Println("init thrift string2i64kv client ")
}

//GetS2I64BinaryClient client by host:port
func GetS2I64BinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := s2i64BinMapPool.Get(aHost, aPort).Get()
	return client
}

//GetS2I64CompactClient get compact client by host:port
func GetS2I64CompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := s2I64CompactMapPool.Get(aHost, aPort).Get()
	return client
}
