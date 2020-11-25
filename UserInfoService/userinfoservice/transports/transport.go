package transports

import (
	"log"

	"github.com/OpenStars/BackendService/UserInfoService/userinfoservice/thrift/gen-go/openstars/userinfoservice"
	"github.com/OpenStars/BackendService/thriftpool"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mUserInfoServiceBinaryMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (userinfoservice.NewTUserInfoServiceClient(c)) }),
		thriftpool.DefaultClose)

	mUserInfoServiceCommpactMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (userinfoservice.NewTUserInfoServiceClient(c)) }),
		thriftpool.DefaultClose)
)

func init() {
	log.Println("init thrift TPubProfileService client ")
}

//GetPubProfileServiceBinaryClient client by host:port
func GetUserInfoServiceBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, err := mUserInfoServiceBinaryMapPool.Get(aHost, aPort).Get()
	if err != nil {
		log.Println("GetPubProfileServiceBinaryClient err", err)
	}
	return client
}

//GetPubProfileServiceCompactClient get compact client by host:port
func GetUserInfoServiceCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mUserInfoServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
