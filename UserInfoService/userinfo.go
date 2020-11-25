package UserInfoService

import (
	"context"
	"errors"

	"github.com/OpenStars/BackendService/UserInfoService/userinfoservice/thrift/gen-go/openstars/userinfoservice"
	"github.com/OpenStars/BackendService/UserInfoService/userinfoservice/transports"
)

type userinforservice struct {
	host string
	port string
	sid  string
}

func (m *userinforservice) GetData(key int64) (*userinfoservice.TUserInfo, error) {

	client := transports.GetUserInfoServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*userinfoservice.TUserInfoServiceClient).GetData(context.Background(), userinfoservice.TUID(key))

	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r.Data == nil {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != userinfoservice.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}
	return r.Data, nil
}
func (m *userinforservice) GetMultiData(keys []int64) (map[userinfoservice.TUID]*userinfoservice.TUserInfo, error) {

	client := transports.GetUserInfoServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*userinfoservice.TUserInfoServiceClient).GetMultiData(context.Background(), keys)

	if err != nil || r == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return r, nil
	// if r.Data == nil {
	// 	return nil, errors.New("Backend service:" + m.sid + " key not found")
	// }
	// if r.ErrorCode != userinfoservice.TErrorCode_EGood {
	// 	return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	// }
	// return r.Data, nil
}
func (m *userinforservice) PutData(uid int64, data *userinfoservice.TUserInfo) error {

	client := transports.GetUserInfoServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*userinfoservice.TUserInfoServiceClient).PutData(context.Background(), userinfoservice.TUID(uid), data)

	if err != nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	// if r.Data == nil {
	// 	return nil, errors.New("Backend service:" + m.sid + " key not found")
	// }
	if r != userinfoservice.TErrorCode_EGood {
		return errors.New("Backend service:" + m.sid + " err:" + r.String())
	}
	return nil
}
