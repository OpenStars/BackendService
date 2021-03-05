package KVStorageService

import (
	"context"
	"errors"

	"github.com/OpenStars/BackendService/KVStorageService/kvstorage/thrift/gen-go/OpenStars/Platform/KVStorage"
	"github.com/OpenStars/BackendService/KVStorageService/kvstorage/transports"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type kvstorageservice struct {
	host string
	port string
	sid  string

	// etcdManager *EndpointsManager.EtcdBackendEndpointManager

	bot_token  string
	bot_chatID int64
	botClient  *tgbotapi.BotAPI
}

func (m *kvstorageservice) notifyEndpointError() {
	if m.botClient != nil {
		msg := tgbotapi.NewMessage(m.bot_chatID, "Hệ thống kiểm soát phát hiện service kvstorageservice có địa chỉ "+m.host+":"+m.port+" đang không hoạt động")
		m.botClient.Send(msg)
	}
}

func (m *kvstorageservice) GetData(key string) (string, error) {
	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }
	client := transports.GetKVStorageCompactClient(m.host, m.port)

	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return "", errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStorage.KVStorageServiceClient).GetData(context.Background(), key)
	if err != nil {
		go m.notifyEndpointError()
		return "", errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if r.ErrorCode != KVStorage.TErrorCode_EGood {
		return "", nil
	}
	return r.Data.Value, nil
}

func (m *kvstorageservice) PutData(key string, value string) (bool, error) {
	client := transports.GetKVStorageCompactClient(m.host, m.port)

	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStorage.KVStorageServiceClient).PutData(context.Background(), key, &KVStorage.KVItem{
		Key:   key,
		Value: value,
	})
	if err != nil {
		go m.notifyEndpointError()
		return false, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if r != KVStorage.TErrorCode_EGood {
		return false, nil
	}
	return true, nil
}

func (m *kvstorageservice) RemoveData(key string) (bool, error) {
	client := transports.GetKVStorageCompactClient(m.host, m.port)

	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStorage.KVStorageServiceClient).RemoveData(context.Background(), key)
	if err != nil {
		go m.notifyEndpointError()
		return false, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if r != KVStorage.TErrorCode_EGood {
		return false, nil
	}
	return true, nil
}

func (m *kvstorageservice) GetListData(keys []string) (results map[string]string, missingkeys []string, err error) {
	client := transports.GetKVStorageCompactClient(m.host, m.port)

	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return nil, nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*KVStorage.KVStorageServiceClient).GetListData(context.Background(), keys)
	if err != nil {
		go m.notifyEndpointError()
		return nil, nil, errors.New("KVCounterService: " + m.sid + " error: " + err.Error())
	}
	defer client.BackToPool()
	if r.ErrorCode != KVStorage.TErrorCode_EGood {
		return nil, nil, nil
	}
	results = make(map[string]string)
	for _, item := range r.Data {
		results[item.Key] = item.Value
	}
	return results, r.Missingkeys, nil
}

func NewClient(sid string, host string, port string) Client {
	client := transports.GetKVStorageCompactClient(host, port)
	if client == nil || client.Client == nil {
		return nil
	}
	kvstorage := &kvstorageservice{
		host: host,
		port: port,
		sid:  sid,

		bot_chatID: -1001469468779,
		bot_token:  "1108341214:AAEKNbFf6PO7Y6UJGK-xepDDOGKlBU2QVCg",
		botClient:  nil,
	}
	bot, err := tgbotapi.NewBotAPI(kvstorage.bot_token)
	if err == nil {
		kvstorage.botClient = bot
	}
	return kvstorage
}
