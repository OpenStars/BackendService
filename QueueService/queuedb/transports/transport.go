package transports

import (
	"fmt"
	"time"

	"github.com/OpenStars/BackendService/QueueService/queuedb/thrift/gen-go/Database/QueueDb"
	telenotification "github.com/OpenStars/BackendService/TeleNotification"
	thriftpool "github.com/OpenStars/BackendService/thriftpoolv2"
	"github.com/apache/thrift/lib/go/thrift"
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
	client := QueueDb.NewQueueDbServiceClientFactory(tf, protocolFactory)

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

var bsGenericMapPool = thriftpool.NewMapPool(1000, 5, 3600, dial, close)

func GetInt2ZsetCompactClient(host, port string) *thriftpool.IdleClient {
	p := bsGenericMapPool.Get(host, port)
	client, err := p.Get()
	if err != nil {
		telenotification.NotifyServiceError(fmt.Sprint("QueueDb totalIdle", p.TotalIdleConn(), "totalConn", p.TotalConn()), host, port, err)
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
func ServiceDisconnect(c *thriftpool.IdleClient, err error) {
	if c == nil {
		return
	}
	bsGenericMapPool.Release(c.Host, c.Port)
	telenotification.NotifyServiceError(c.SID, c.Host, c.Port, err)
}
