package transports

import (
	"github.com/OpenStars/BackendService/Int2StringService/i2skv/thrift/gen-go/OpenStars/Common/I2SKV"
	"github.com/OpenStars/BackendService/thriftpool"
	"github.com/apache/thrift/lib/go/thrift"

	// "github.com/OpenStars/backendclients/go//gen-go/OpenStars/Common/I2SKV" //Todo: Fix this
	"fmt"
)

var (
	mTI2StringServiceBinaryMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (I2SKV.NewTI2StringServiceClient(c)) }),
		thriftpool.DefaultClose)

	mTI2StringServiceCommpactMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (I2SKV.NewTI2StringServiceClient(c)) }),
		thriftpool.DefaultClose)
)

func init() {
	fmt.Println("init thrift TI2StringService client ")
}

//GetTI2StringServiceBinaryClient client by host:port
func GetTI2StringServiceBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mTI2StringServiceBinaryMapPool.Get(aHost, aPort).Get()
	return client
}

//GetTI2StringServiceCompactClient get compact client by host:port
func GetTI2StringServiceCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mTI2StringServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
