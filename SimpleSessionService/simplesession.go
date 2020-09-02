package SimpleSessionService

import (
	"context"
	"errors"
	"log"

	"github.com/OpenStars/BackendService/EndpointsManager"
	"github.com/OpenStars/BackendService/SimpleSessionService/simplesession/thrift/gen-go/simplesession"
	"github.com/OpenStars/BackendService/SimpleSessionService/simplesession/transports"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type simpleSessionClient struct {
	host string
	port string
	sid  string

	etcdManager *EndpointsManager.EtcdBackendEndpointManager

	bot_token  string
	bot_chatID int64
	botClient  *tgbotapi.BotAPI
}

func (m *simpleSessionClient) notifyEndpointError() {
	if m.botClient != nil {
		msg := tgbotapi.NewMessage(m.bot_chatID, "Hệ thống kiểm soát endpoint phát hiện endpoint sid "+m.sid+" address "+m.host+":"+m.port+" đang không hoạt động")
		m.botClient.Send(msg)
	}
}

func (m *simpleSessionClient) GetSession(sskey string) (*simplesession.TUserSessionInfo, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetSimpleSessionCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*simplesession.TSimpleSessionService_WClient).GetSession(context.Background(), simplesession.TSessionKey(sskey))
	if err != nil {
		go m.notifyEndpointError()
		return nil, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r == nil || r.Error == simplesession.TErrorCode_EFailed || r.UserInfo == nil {
		return nil, nil
	}
	return r.UserInfo, nil
}

func (m *simpleSessionClient) CreateSession(uid int64, deviceInfo string, data string, expiredTime int64) (string, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetSimpleSessionCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return "", errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	userinfo := &simplesession.TUserSessionInfo{
		Version:     1,
		Code:        simplesession.TSessionCode_EFullRight,
		UID:         simplesession.TUID(uid),
		DeviceInfo:  deviceInfo,
		Data:        data,
		ExpiredTime: expiredTime,
	}
	r, err := client.Client.(*simplesession.TSimpleSessionService_WClient).CreateSession(context.Background(), userinfo)
	if err != nil {
		go m.notifyEndpointError()
		return "", errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	if r.Error == simplesession.TErrorCode_EFailed || r == nil || r.Session == nil {
		return "", nil
	}
	return string(*r.Session), nil
}

func (m *simpleSessionClient) RemoveSession(sskey string) (bool, error) {
	if m.etcdManager != nil {
		h, p, err := m.etcdManager.GetEndpoint(m.sid)
		if err != nil {
			log.Println("EtcdManager get endpoints", "err", err)
		} else {
			m.host = h
			m.port = p
		}
	}
	client := transports.GetSimpleSessionCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return false, errors.New("Can not connect to backend service: " + m.sid + "host: " + m.host + "port: " + m.port)
	}

	r, err := client.Client.(*simplesession.TSimpleSessionService_WClient).RemoveSession(context.Background(), simplesession.TSessionKey(sskey))
	if err != nil {
		go m.notifyEndpointError()
		return false, errors.New("Backend service:" + m.sid + " err:" + err.Error())
	}
	defer client.BackToPool()
	return r, nil
}
