package String2Int64Service

import (
	"context"
	"errors"

	"github.com/OpenStars/BackendService/String2Int64Service/s2i64kv/thrift/gen-go/OpenStars/Common/S2I64KV"
	"github.com/OpenStars/BackendService/String2Int64Service/s2i64kv/transports"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type String2Int64Service struct {
	host string
	port string
	sid  string

	// etcdManager *EndpointsManager.EtcdBackendEndpointManager

	bot_token  string
	bot_chatID int64
	botClient  *tgbotapi.BotAPI
}

func (m *String2Int64Service) notifyEndpointError() {
	if m.botClient != nil {
		msg := tgbotapi.NewMessage(m.bot_chatID, "Hệ thống kiểm soát phát hiện service int2string có địa chỉ "+m.host+":"+m.port+" đang không hoạt động")
		m.botClient.Send(msg)
	}
}
func (m *String2Int64Service) PutData(key string, value int64) error {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetS2I64CompactClient(m.host, m.port)

	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return errors.New("Can not connect to model")
	}

	tkey := S2I64KV.TKey(key)
	tvalue := &S2I64KV.TI64Value{
		Value: value,
	}

	_, err := client.Client.(*S2I64KV.TString2I64KVServiceClient).PutData(context.Background(), tkey, tvalue)

	if err != nil {
		go m.notifyEndpointError()
		return errors.New("String2Int64Service sid: " + m.sid + " address: " + m.host + ":" + m.port + " err: " + err.Error())
	}
	defer client.BackToPool()
	return nil
}

func (m *String2Int64Service) GetData(key string) (int64, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetS2I64CompactClient(m.host, m.port)

	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return -1, errors.New("Can not connect to model")
	}

	tkey := S2I64KV.TKey(key)
	r, err := client.Client.(*S2I64KV.TString2I64KVServiceClient).GetData(context.Background(), tkey)

	if err != nil {
		go m.notifyEndpointError()
		return -1, errors.New("String2Int64Service sid: " + m.sid + " address: " + m.host + ":" + m.port + " err: " + err.Error())
	}
	defer client.BackToPool()
	if r == nil || r.Data == nil || r.ErrorCode != S2I64KV.TErrorCode_EGood || r.Data.Value <= 0 {

		return -1, errors.New("Can not found key")
	}
	return r.Data.Value, nil
}

func (m *String2Int64Service) CasData(key string, value int64) (sucess bool, oldvalue int64, err error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetS2I64CompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return false, -1, errors.New("Can not connect to model")
	}

	var aCas = &S2I64KV.TCasValue{OldValue: 0, NewValue_: value}
	r, err := client.Client.(*S2I64KV.TString2I64KVServiceClient).CasData(context.Background(), S2I64KV.TKey(key), aCas)
	if err != nil {
		go m.notifyEndpointError()
		return false, -1, errors.New("String2Int64Service sid: " + m.sid + " address: " + m.host + ":" + m.port + " err: " + err.Error())
	}
	defer client.BackToPool()

	if r != nil && r.GetOldValue() == 0 {
		return true, value, nil
	}
	return false, r.GetOldValue(), nil
}

func NewClient(etcdServers []string, sid, host, port string) Client {
	s2isv := &String2Int64Service{
		host: host,
		port: port,
		sid:  sid,
		// etcdManager: EndpointsManager.GetEtcdBackendEndpointManagerSingleton(etcdServers),
		bot_chatID: -1001469468779,
		bot_token:  "1108341214:AAEKNbFf6PO7Y6UJGK-xepDDOGKlBU2QVCg",
		botClient:  nil,
	}
	bot, err := tgbotapi.NewBotAPI(s2isv.bot_token)
	if err == nil {
		s2isv.botClient = bot
	}
	// if s2isv.etcdManager == nil {

	// 	return s2isv
	// }
	// err = s2isv.etcdManager.SetDefaultEntpoint(sid, host, port)
	// if err != nil {
	// 	log.Println("SetDefaultEndpoint sid", sid, "err", err)
	// 	return s2isv
	// }
	// s2isv.etcdManager.GetAllEndpoint(serviceID)
	return s2isv
}
