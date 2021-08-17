package PublicProfileService

import (
	"context"
	"errors"

	"github.com/OpenStars/BackendService/PublicProfileService/tpubprofileservice/thrift/gen-go/openstars/pubprofile"
	transports "github.com/OpenStars/BackendService/PublicProfileService/tpubprofileservice/transportsv2"
)

type pubprofileclient struct {
	host string
	port string
}

func (m *pubprofileclient) GetProfileByUID(uid int64) (r *pubprofile.ProfileData, err error) {
	client := transports.GetPubProfileServiceBinaryClient(m.host, m.port)
	if client == nil || client.Client == nil {

		transports.ServiceDisconnect(client)
		return nil, errors.New("Can not connect to backend service host: " + m.host + " port: " + m.port)
	}
	resp, err := client.Client.(*pubprofile.PubProfileServiceClient).GetProfileByUID(context.Background(), uid)
	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("Backend service err:" + err.Error())
	}

	defer transports.BackToPool(client)

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
		transports.ServiceDisconnect(client)
		return nil, errors.New("Can not connect to backend service host: " + m.host + " port: " + m.port)
	}

	resp, err := client.Client.(*pubprofile.PubProfileServiceClient).GetProfileByPubkey(context.Background(), pubkey)
	if err != nil {
		transports.ServiceDisconnect(client)
		return nil, errors.New("Backend service err:" + err.Error())
	}

	defer transports.BackToPool(client)

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
	defer transports.BackToPool(client)

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
	defer transports.BackToPool(client)

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
	defer transports.BackToPool(client)

	if resp != nil {
		return resp.Resp, nil
	}
	return false, nil
}
