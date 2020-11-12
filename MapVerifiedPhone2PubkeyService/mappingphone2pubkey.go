package MapVerifiedPhone2PubkeyService

import (
	"context"
	"errors"

	"github.com/OpenStars/BackendService/MapVerifiedPhone2PubkeyService/mapphone2pubkey/thrift/gen-go/OpenStars/Common/MapPhoneNumberPubkeyKV"
	"github.com/OpenStars/BackendService/MapVerifiedPhone2PubkeyService/mapphone2pubkey/transports"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type MappingPhone2PubkeyServiceModel struct {
	host string
	port string
	sid  string
	// etcdManager *EndpointsManager.EtcdBackendEndpointManager

	bot_token  string
	bot_chatID int64
	botClient  *tgbotapi.BotAPI
}

func (m *MappingPhone2PubkeyServiceModel) PutData(pubkey string, phonenumber string) (bool, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetTMapPhoneNumberPubkeyKVServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return false, errors.New("Can not connect to model")
	}

	r, err := client.Client.(*MapPhoneNumberPubkeyKV.TMapPhoneNumberPubkeyKVServiceClient).PutData(context.Background(), pubkey, phonenumber)
	if err != nil {
		go m.notifyEndpointError()
		return false, errors.New("Serviceid:" + m.sid + " address:" + m.host + ":" + m.port + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r != MapPhoneNumberPubkeyKV.TErrorCode_EGood {
		return false, nil
	}
	return true, nil
}

func (m *MappingPhone2PubkeyServiceModel) GetPhoneNumberByPubkey(pubkey string) (string, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetTMapPhoneNumberPubkeyKVServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return "", errors.New("Can not connect to model")
	}

	r, err := client.Client.(*MapPhoneNumberPubkeyKV.TMapPhoneNumberPubkeyKVServiceClient).GetPhoneNumberByPubkey(context.Background(), pubkey)
	if err != nil {
		go m.notifyEndpointError()
		return "", errors.New("Serviceid:" + m.sid + " address:" + m.host + ":" + m.port + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r.ErrorCode != MapPhoneNumberPubkeyKV.TErrorCode_EGood {
		return "", nil
	}
	return r.Data.Value, nil
}

func (m *MappingPhone2PubkeyServiceModel) GetPubkeyByPhoneNumber(phonenumber string) (string, error) {

	// if m.etcdManager != nil {
	// 	h, p, err := m.etcdManager.GetEndpoint(m.sid)
	// 	if err != nil {
	// 		log.Println("EtcdManager get endpoints", "err", err)
	// 	} else {
	// 		m.host = h
	// 		m.port = p
	// 	}
	// }

	client := transports.GetTMapPhoneNumberPubkeyKVServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return "", errors.New("Can not connect to model")
	}

	r, err := client.Client.(*MapPhoneNumberPubkeyKV.TMapPhoneNumberPubkeyKVServiceClient).GetPubkeyByPhoneNumber(context.Background(), phonenumber)
	if err != nil {
		go m.notifyEndpointError()
		return "", errors.New("Serviceid:" + m.sid + " address:" + m.host + ":" + m.port + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r.ErrorCode != MapPhoneNumberPubkeyKV.TErrorCode_EGood {
		return "", nil
	}
	return r.Data.Value, nil
}

func (m *MappingPhone2PubkeyServiceModel) notifyEndpointError() {
	if m.botClient != nil {
		msg := tgbotapi.NewMessage(m.bot_chatID, "Hệ thống kiểm soát endpoint phát hiện endpoint sid "+m.sid+" address "+m.host+":"+m.port+" đang không hoạt động")
		m.botClient.Send(msg)
	}

}

func NewClient(etcdServer []string, sid, host, port string) Client {
	mapphone2pub := &MappingPhone2PubkeyServiceModel{
		host: host,
		port: port,
		sid:  sid,
		// etcdManager: EndpointsManager.GetEtcdBackendEndpointManagerSingleton(etcdServer),
		bot_chatID: -1001469468779,
		bot_token:  "1108341214:AAEKNbFf6PO7Y6UJGK-xepDDOGKlBU2QVCg",
		botClient:  nil,
	}
	bot, err := tgbotapi.NewBotAPI(mapphone2pub.bot_token)
	if err == nil {
		mapphone2pub.botClient = bot
	}
	// if mapphone2pub.etcdManager == nil {
	// 	return mapphone2pub
	// }
	// err = mapphone2pub.etcdManager.SetDefaultEntpoint(sid, host, port)
	// if err != nil {
	// 	log.Println("SetDefaultEndpoint sid", sid, "err", err)
	// 	return nil
	// }

	return mapphone2pub
}
