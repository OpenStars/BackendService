package transportsv2

import (
	"errors"
	"fmt"
	"time"

	telenotification "github.com/OpenStars/BackendService/TeleNotification"
	thriftpool "github.com/OpenStars/BackendService/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"

	"github.com/OpenStars/BackendService/KVStorageService/kvstorage/thrift/gen-go/OpenStars/Platform/KVStorage"
)

func dial(addr, port string, connTimeout time.Duration) (*thriftpool.IdleClient, error) {

	socket, err := thrift.NewTSocketTimeout(fmt.Sprintf("%s:%s", addr, port), connTimeout)
	if err != nil {
		return nil, err
	}
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTBufferedTransportFactory(10000))
	protocolFactory := thrift.NewTCompactProtocolFactory()
	tf, err := transportFactory.GetTransport(socket)
	if err != nil {
		return nil, err
	}
	client := KVStorage.NewKVStorageServiceClientFactory(tf, protocolFactory)

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
		Host:   addr,
		Port:   port,
	}, nil
}

func close(c *thriftpool.IdleClient) error {
	err := c.Socket.Close()
	//err = c.Client.(*tutorial.PlusServiceClient).Transport.Close()
	return err
}

var bsGenericMapPool = thriftpool.NewMapPool(4, 5, 3600, dial, close)

func GetKVStorageCompactClient(host, port string) *thriftpool.IdleClient {
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

	bsGenericMapPool.Get(c.Host, c.Port).Put(c)
}
func ServiceDisconnect(c *thriftpool.IdleClient) {
	if c == nil {
		return
	}
	bsGenericMapPool.Release(c.Host, c.Port)
	telenotification.NotifyServiceError("", c.Host, c.Port, errors.New("service disconnect"))
}
