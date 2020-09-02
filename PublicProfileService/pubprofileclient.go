package PubProfileClient

import (
	"context"
	"errors"

	"github.com/OpenStars/BackendService/PublicProfileService/tpubprofileservice/thrift/gen-go/openstars/pubprofile"
	"github.com/OpenStars/BackendService/PublicProfileService/tpubprofileservice/transports"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type pubprofileclient struct {
	host string
	port string

	bot_token  string
	bot_chatID int64
	botClient  *tgbotapi.BotAPI
}

func (m *pubprofileclient) notifyEndpointError() {
	if m.botClient != nil {
		msg := tgbotapi.NewMessage(m.bot_chatID, "Hệ thống kiểm soát endpoint phát hiện pubprofile endpoint address "+m.host+":"+m.port+" đang không hoạt động")
		m.botClient.Send(msg)
	}
}

func (m *pubprofileclient) GetProfileByUID(uid int64) (r *pubprofile.ProfileData, err error) {
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
		return resp.ProfileData, nil
	}
	return nil, errors.New("Get data nil")
}

func (m *pubprofileclient) GetProfileByPubkey(pubkey string) (r *pubprofile.ProfileData, err error) {
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
		return resp.Resp, nil
	}
	return false, nil
}
