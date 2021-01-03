package transports

import (
	"log"

	"github.com/OpenStars/BackendService/MoneyStorageService/money/moneyservice"
	"github.com/OpenStars/BackendService/thriftpool"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mMoneyStorageServiceBinaryMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (moneyservice.NewTMoneyAgentServiceClient(c)) }),
		thriftpool.DefaultClose)

	mMoneyStorageServiceCommpactMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (moneyservice.NewTMoneyAgentServiceClient(c)) }),
		thriftpool.DefaultClose)
)

//GetPubProfileServiceBinaryClient client by host:port
func GetMoneyStorageServiceBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, err := mMoneyStorageServiceBinaryMapPool.Get(aHost, aPort).Get()
	if err != nil {
		log.Println("GetPubProfileServiceBinaryClient err", err)
	}
	return client
}

//GetPubProfileServiceCompactClient get compact client by host:port
func GetMoneyStorageServiceCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mMoneyStorageServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
