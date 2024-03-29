package transportsv2

import (
	"errors"
	"fmt"
	"strings"
	"time"

	telenotification "github.com/OpenStars/BackendService/TeleNotification"
	thriftpool "github.com/OpenStars/BackendService/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"

	"github.com/OpenStars/BackendService/PublicProfileService/tpubprofileservice/thrift/gen-go/openstars/pubprofile"
)

func dial(addr, port string, connTimeout time.Duration) (*thriftpool.IdleClient, error) {

	socket, err := thrift.NewTSocketTimeout(fmt.Sprintf("%s:%s", addr, port), connTimeout)
	if err != nil {
		return nil, err
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTBufferedTransportFactory(10000))
	protocolFactory := thrift.NewTBinaryProtocolFactory(true, true)
	tf, err := transportFactory.GetTransport(socket)
	if err != nil {
		return nil, err
	}
	client := pubprofile.NewPubProfileServiceClientFactory(tf, protocolFactory)

	if err != nil {
		return nil, err
	}
	err = tf.Open()
	if err != nil {
		return nil, err
	}
	return &thriftpool.IdleClient{
		Client: client,
		Socket: socket,
	}, nil
}

func close(c *thriftpool.IdleClient) error {
	err := c.Socket.Close()
	//err = c.Client.(*tutorial.PlusServiceClient).Transport.Close()
	return err
}

var bsGenericMapPool = thriftpool.NewMapPool(1000, 5, 3600, dial, close)

func GetPubProfileServiceBinaryClient(host, port string) *thriftpool.IdleClient {
	client, err := bsGenericMapPool.Get(host, port).Get()
	if err != nil {
		telenotification.NotifyServiceError("", host, port, err)
		return nil
	}
	return client
}

func BackToPool(c *thriftpool.IdleClient) {
	if c == nil {
		return
	}
	netarr := strings.Split(c.Socket.Addr().String(), ":")
	bsGenericMapPool.Get(netarr[0], netarr[1]).Put(c)
}
func ServiceDisconnect(c *thriftpool.IdleClient) {
	if c == nil {
		return
	}
	netarr := strings.Split(c.Socket.Addr().String(), ":")
	bsGenericMapPool.Release(netarr[0], netarr[1])
	telenotification.NotifyServiceError("", netarr[0], netarr[1], errors.New("service disconnect"))
}
