package transports

import (
	"log"

	"github.com/OpenStars/BackendService/PublicProfileService/tpubprofileservice/thrift/gen-go/openstars/pubprofile"
	"github.com/OpenStars/BackendService/thriftpool"
	"github.com/apache/thrift/lib/go/thrift"
)

var (
	mPubProfileServiceBinaryMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFunc(func(c thrift.TClient) interface{} { return (pubprofile.NewPubProfileServiceClient(c)) }),
		thriftpool.DefaultClose)

	mPubProfileServiceCommpactMapPool = thriftpool.NewMapPool(1000, 3600, 3600,
		thriftpool.GetThriftClientCreatorFuncCompactProtocol(func(c thrift.TClient) interface{} { return (pubprofile.NewPubProfileServiceClient(c)) }),
		thriftpool.DefaultClose)
)

func init() {
	log.Println("init thrift TPubProfileService client ")
}

//GetPubProfileServiceBinaryClient client by host:port
func GetPubProfileServiceBinaryClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, err := mPubProfileServiceBinaryMapPool.Get(aHost, aPort).Get()
	if err != nil {
		log.Println("GetPubProfileServiceBinaryClient err", err)
	}
	return client
}

//GetPubProfileServiceCompactClient get compact client by host:port
func GetPubProfileServiceCompactClient(aHost, aPort string) *thriftpool.ThriftSocketClient {
	client, _ := mPubProfileServiceCommpactMapPool.Get(aHost, aPort).Get()
	return client
}
