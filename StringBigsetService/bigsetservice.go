package StringBigsetService

import (
	"context"
	"errors"
	"log"
	"sync"
	"time"

	"github.com/OpenStars/BackendService/StringBigsetService/bigset/thrift/gen-go/openstars/core/bigset/generic"
	"github.com/OpenStars/BackendService/StringBigsetService/bigset/transports"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var reconnect = true
var mureconnect sync.Mutex

type StringBigsetService struct {
	host string
	port string
	sid  string

	// etcdManager *EndpointsManager.EtcdBackendEndpointManager

	bot_token          string
	bot_chatID         int64
	botClient          *tgbotapi.BotAPI
	agentCallerName    string
	agentCallerAddress string
}

func (m *StringBigsetService) SetAgentCaller(agentName string, agentAddress string) {
	m.agentCallerAddress = agentAddress
	m.agentCallerName = agentName
}

func (m *StringBigsetService) notifyEndpointError(err error) {
	if m.botClient != nil {
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		msg := tgbotapi.NewMessage(m.bot_chatID, "Hệ thống kiểm soát endpoint phát hiện endpoint sid "+m.sid+" address "+m.host+":"+m.port+" đang không hoạt động"+" từ agent caller "+m.agentCallerName+"địa chỉ "+m.agentCallerAddress+" error "+errString)
		m.botClient.Send(msg)
	}

}

func (m *StringBigsetService) handleEventChangeEndpoint(host, port string) {
	m.host = host
	m.port = port
	log.Println("Change config endpoint serviceID", m.sid, m.host, ":", m.port)
}

func NewClient(etcdEndpoints []string, sid string, defaultEndpointsHost string, defaultEndpointPort string) Client {
	log.Println("Init StringBigset Service sid", sid, "address", defaultEndpointsHost+":"+defaultEndpointPort)
	client := transports.GetBsGenericClient(defaultEndpointsHost, defaultEndpointPort)
	if client == nil || client.Client == nil {
		return nil
	}
	stringbs := &StringBigsetService{
		host: defaultEndpointsHost,
		port: defaultEndpointPort,
		sid:  sid,
		// etcdManager: EndpointsManager.GetEtcdBackendEndpointManagerSingleton(etcdEndpoints),
		bot_chatID: -1001469468779,
		bot_token:  "1108341214:AAEKNbFf6PO7Y6UJGK-xepDDOGKlBU2QVCg",
		botClient:  nil,
	}
	bot, err := tgbotapi.NewBotAPI(stringbs.bot_token)
	if err == nil {
		stringbs.botClient = bot
	}
	// if stringbs.etcdManager == nil {
	// 	return stringbs
	// }
	// err = stringbs.etcdManager.SetDefaultEntpoint(sid, defaultEndpointsHost, defaultEndpointPort)
	// if err != nil {
	// 	log.Println("SetDefaultEndpoint sid", sid, "err", err)
	// 	return nil
	// }
	return stringbs
}

func NewClientWithMonitor(etcdEndpoints []string, sid string, host string, port string, bot_token string, bot_chatID int64) Client {
	// 1108341214:AAEKNbFf6PO7Y6UJGK-xepDDOGKlBU2QVCg
	// -1001469468779

	log.Println("Init StringBigset Service sid", sid, "address", host+":"+port)
	stringbs := &StringBigsetService{
		host: host,
		port: port,
		sid:  sid,
		// etcdManager: EndpointsManager.GetEtcdBackendEndpointManagerSingleton(etcdEndpoints),
		botClient:  nil,
		bot_chatID: bot_chatID,
		bot_token:  bot_token,
	}
	bot, err := tgbotapi.NewBotAPI(bot_token)
	if err == nil {
		stringbs.botClient = bot
	}
	// if stringbs.etcdManager == nil {
	// 	return stringbs
	// }
	// err = stringbs.etcdManager.SetDefaultEntpoint(sid, host, port)
	// if err != nil {
	// 	log.Println("SetDefaultEndpoint sid", sid, "err", err)
	// 	return nil
	// }
	return stringbs
}

func (m *StringBigsetService) TotalStringKeyCount() (r int64, err error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)

	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return 0, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err = client.Client.(*generic.TStringBigSetKVServiceClient).TotalStringKeyCount(ctx)

	if err != nil {
		go m.notifyEndpointError(err)
		return 0, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	return r, nil
}

func (m *StringBigsetService) GetListKey(fromIndex int64, count int32) ([]string, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetListKey(ctx, fromIndex, count)

	if err != nil {
		go m.notifyEndpointError(err)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	var listKey []string
	for _, item := range r {
		listKey = append(listKey, string(item))
	}
	return listKey, nil
}

func (m *StringBigsetService) BsPutItem(bskey string, itemKey, itemVal string) (bool, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	// log.Println("BsPutItem host", m.host+":"+m.port)
	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsPutItem(ctx, generic.TStringKey(bskey), &generic.TItem{
		Key:   []byte(itemKey),
		Value: []byte(itemVal),
	})

	if err != nil {
		go m.notifyEndpointError(err)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return false, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if r.Error != generic.TErrorCode_EGood {
		return false, nil
	}
	// if r.Ok == false {
	// 	err := errors.New("Can not write bskey: " + bskey + " itemkey: " + itemKey)
	// 	go m.notifyEndpointError(err)
	// 	return false, err
	// }
	return true, nil

}

func (m *StringBigsetService) BsRangeQuery(bskey string, startKey string, endKey string) ([]*generic.TItem, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRangeQuery(ctx, generic.TStringKey(bskey), generic.TItemKey(startKey), generic.TItemKey(endKey))
	if err != nil {
		go m.notifyEndpointError(err)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}

	return rs.Items.Items, nil
}

// BsRangeQueryByPage get >= startkey && <= endkey có chia page theo begin and end
func (m *StringBigsetService) BsRangeQueryByPage(bskey string, startKey, endKey string, begin, end int64) ([]*generic.TItem, int64, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)

	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return nil, -1, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRangeQuery(ctx, generic.TStringKey(bskey), generic.TItemKey(startKey), generic.TItemKey(endKey))
	if err != nil {
		go m.notifyEndpointError(err)
		return nil, -1, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	if r.Items != nil && r.Items.Items != nil && len(r.Items.Items) > 0 { // pagination
		if begin < 0 {
			begin = 0
		}
		if end > int64(len(r.Items.Items)) {
			end = int64(len(r.Items.Items))
		}
		total := int64(len(r.Items.Items))
		r.Items.Items = r.Items.Items[begin:end]

		return r.Items.Items, total, nil
	}

	return nil, 0, nil
}

func (m *StringBigsetService) BsGetItem(bskey string, itemkey string) (*generic.TItem, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	// fmt.Printf("[BsGetItem] get client host = %s, %s, key = %s, %s \n", m.host, m.port, bskey, itemkey)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetItem(ctx, generic.TStringKey(bskey), generic.TItemKey(itemkey))
	if err != nil {
		go m.notifyEndpointError(err)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	if r.Error != generic.TErrorCode_EGood || r.Item == nil || r.Item.Key == nil {
		return nil, nil
	}
	return r.Item, nil
}

func (m *StringBigsetService) GetTotalCount(bskey string) (int64, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return 0, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	r, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetTotalCount(ctx, generic.TStringKey(bskey))

	if err != nil {
		go m.notifyEndpointError(err)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return 0, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()

	if r <= 0 {
		return 0, nil
	}
	return r, nil

}

func (m *StringBigsetService) GetBigSetInfoByName(bskey string) (*generic.TStringBigSetInfo, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).GetBigSetInfoByName(ctx, generic.TStringKey(bskey))
	if err != nil {
		go m.notifyEndpointError(err)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()

	if rs.Info == nil {
		return nil, nil
	}
	return rs.Info, nil

}

func (m *StringBigsetService) RemoveAll(bskey string) (bool, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := client.Client.(*generic.TStringBigSetKVServiceClient).RemoveAll(ctx, generic.TStringKey(bskey))
	if err != nil {
		go m.notifyEndpointError(err)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()
	return true, nil
}
func (m *StringBigsetService) CreateStringBigSet(bskey string) (*generic.TStringBigSetInfo, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).CreateStringBigSet(ctx, generic.TStringKey(bskey))
	if err != nil {
		go m.notifyEndpointError(err)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	defer client.BackToPool()

	if rs.Info == nil {
		return nil, nil
	}
	return rs.Info, nil
}

func (m *StringBigsetService) BsGetSlice(bskey string, fromPos int32, count int32) ([]*generic.TItem, error) {
	if count == 0 {
		return nil, nil
	}
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSlice(ctx, generic.TStringKey(bskey), fromPos, count)
	if err != nil {
		go m.notifyEndpointError(err)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}

	return rs.Items.Items, nil
}

func (m *StringBigsetService) BsGetSliceR(bskey string, fromPos int32, count int32) ([]*generic.TItem, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceR(ctx, generic.TStringKey(bskey), fromPos, count)
	if err != nil {
		go m.notifyEndpointError(err)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}
	return rs.Items.Items, nil
}

func (m *StringBigsetService) BsRemoveItem(bskey string, itemkey string) (bool, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	ok, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsRemoveItem(ctx, generic.TStringKey(bskey), generic.TItemKey(itemkey))
	if err != nil {
		go m.notifyEndpointError(err)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return false, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	return ok, nil
}

func (m *StringBigsetService) BsMultiPut(bskey string, lsItems []*generic.TItem) (bool, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	itemset := &generic.TItemSet{
		Items: lsItems,
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsMultiPut(ctx, generic.TStringKey(bskey), itemset, false, false)
	if err != nil {
		go m.notifyEndpointError(err)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return false, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood {
		return false, nil
	}
	return true, nil
}

func (m *StringBigsetService) BsGetSliceFromItem(bskey string, fromKey string, count int32) ([]*generic.TItem, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceFromItem(ctx, generic.TStringKey(bskey), generic.TItemKey(fromKey), count)
	if err != nil {
		go m.notifyEndpointError(err)
		// client = transports.NewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}

	return rs.Items.Items, nil
}

func (m *StringBigsetService) BsGetSliceFromItemR(bskey string, fromKey string, count int32) ([]*generic.TItem, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetBsGenericClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError(errors.New("Client transport null"))
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	rs, err := client.Client.(*generic.TStringBigSetKVServiceClient).BsGetSliceFromItemR(ctx, generic.TStringKey(bskey), generic.TItemKey(fromKey), count)
	if err != nil {
		go m.notifyEndpointError(err)
		// client = transports.NeewGetBsGenericClient(m.host, m.port)
		return nil, errors.New("StringBigsetSerice: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if rs.Error != generic.TErrorCode_EGood || rs.Items == nil || rs.Items.Items == nil {
		return nil, nil
	}
	return rs.Items.Items, nil
}
