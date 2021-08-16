package PublicProfileService

import (
	"context"
	"errors"
	"time"

	"github.com/OpenStars/BackendService/PublicProfileService/tpubprofileservice/thrift/gen-go/openstars/pubprofile"
	"github.com/OpenStars/BackendService/PublicProfileService/tpubprofileservice/transports"
	telenotification "github.com/OpenStars/BackendService/TeleNotification"
	"github.com/bluele/gcache"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type pubprofileclient struct {
	host string
	port string

	bot_token     string
	bot_chatID    int64
	botClient     *tgbotapi.BotAPI
	cache         gcache.Cache
	cacheExperied time.Duration
}

func (m *pubprofileclient) notifyEndpointError() {
	telenotification.NotifyServiceError("pubprofile", m.host, m.port, nil)
	// if m.botClient != nil {
	// 	msg := tgbotapi.NewMessage(m.bot_chatID, "Hệ thống kiểm soát endpoint phát hiện pubprofile endpoint address "+m.host+":"+m.port+" đang không hoạt động")
	// 	m.botClient.Send(msg)
	// }
}

func (m *pubprofileclient) GetProfileByUID(uid int64) (r *pubprofile.ProfileData, err error) {
	if m.cache != nil {
		profileV, _ := m.cache.Get(uid)
		if profileV != nil {
			profile := profileV.(*pubprofile.ProfileData)
			return profile, nil
		}
	}
	client := transports.GetPubProfileServiceBinaryClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service host: " + m.host + " port: " + m.port)
	}

	resp, err := client.Client.(*pubprofile.PubProfileServiceClient).GetProfileByUID(context.Background(), uid)
	if err != nil {
		go m.notifyEndpointError()
		return nil, errors.New("Backend service err:" + err.Error())
	}

	defer client.BackToPool()

	if resp != nil && resp.ProfileData != nil {

		if resp.ProfileData.Pubkey == "" && resp.ProfileData.DisplayName == "" && resp.ProfileData.LastModified == 0 {
			return nil, errors.New("Profile not existed")
		}
		if m.cache != nil {
			m.cache.SetWithExpire(uid, resp.ProfileData, m.cacheExperied)
		}
		return resp.ProfileData, nil
	}
	return nil, errors.New("Get data nil")
}

func (m *pubprofileclient) GetProfileByPubkey(pubkey string) (r *pubprofile.ProfileData, err error) {
	if m.cache != nil {
		profileV, _ := m.cache.Get(pubkey)
		if profileV != nil {
			profile := profileV.(*pubprofile.ProfileData)
			return profile, nil
		}
	}
	client := transports.GetPubProfileServiceBinaryClient(m.host, m.port)
	if client == nil || client.Client == nil {
		go m.notifyEndpointError()
		return nil, errors.New("Can not connect to backend service host: " + m.host + " port: " + m.port)
	}

	resp, err := client.Client.(*pubprofile.PubProfileServiceClient).GetProfileByPubkey(context.Background(), pubkey)
	if err != nil {
		go m.notifyEndpointError()
		return nil, errors.New("Backend service err:" + err.Error())
	}
	defer client.BackToPool()

	if resp != nil && resp.ProfileData != nil {
		if resp.ProfileData.Pubkey == "" && resp.ProfileData.DisplayName == "" && resp.ProfileData.LastModified == 0 {
			return nil, errors.New("Profile not existed")
		}
		if m.cache != nil {
			m.cache.SetWithExpire(pubkey, resp.ProfileData, m.cacheExperied)
		}
		return resp.ProfileData, nil
	}
	return nil, errors.New("Get data nil")
}

func (m *pubprofileclient) UpdateProfileByPubkey(pubkey string, profileUpdate *pubprofile.ProfileData) (r bool, err error) {
	client := transports.GetPubProfileServiceBinaryClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service host: " + m.host + " port: " + m.port)
	}

	resp, err := client.Client.(*pubprofile.PubProfileServiceClient).UpdateProfileByPubkey(context.Background(), pubkey, profileUpdate)
	if err != nil {
		return false, errors.New("Backend service err:" + err.Error())
	}
	defer client.BackToPool()

	if resp != nil {
		if m.cache != nil {
			m.cache.SetWithExpire(pubkey, profileUpdate, m.cacheExperied)
		}
		return resp.Resp, nil
	}
	return false, nil
}

func (m *pubprofileclient) SetProfileByPubkey(pubkey string, profileUpdate *pubprofile.ProfileData) (r bool, err error) {
	client := transports.GetPubProfileServiceBinaryClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service host: " + m.host + " port: " + m.port)
	}

	resp, err := client.Client.(*pubprofile.PubProfileServiceClient).SetProfileByPubkey(context.Background(), pubkey, profileUpdate)
	if err != nil {
		return false, errors.New("Backend service err:" + err.Error())
	}
	defer client.BackToPool()

	if resp != nil {
		if m.cache != nil {
			m.cache.SetWithExpire(pubkey, profileUpdate, m.cacheExperied)
		}
		return resp.Resp, nil
	}
	return false, nil
}

func (m *pubprofileclient) UpdateProfileByUID(uid int64, profileUpdate *pubprofile.ProfileData) (r bool, err error) {
	client := transports.GetPubProfileServiceBinaryClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return false, errors.New("Can not connect to backend service host: " + m.host + " port: " + m.port)
	}

	resp, err := client.Client.(*pubprofile.PubProfileServiceClient).UpdateProfileByUID(context.Background(), uid, profileUpdate)
	if err != nil {
		return false, errors.New("Backend service err:" + err.Error())
	}
	defer client.BackToPool()

	if resp != nil {
		if m.cache != nil {
			m.cache.SetWithExpire(uid, profileUpdate, m.cacheExperied)
		}
		return resp.Resp, nil
	}
	return false, nil
}
