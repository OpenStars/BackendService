package PassportService

import (
	"context"
	"errors"

	"github.com/OpenStars/BackendService/PassportService/ppassport/thrift/gen-go/OpenStars/Platform/Passport"
	"github.com/OpenStars/BackendService/PassportService/ppassport/transports"
	telenotification "github.com/OpenStars/BackendService/TeleNotification"
)

type ppassportservice struct {
	host string
	port string
	sid  string
}

func (m *ppassportservice) GetData(key int64) (*Passport.TPassportInfo, error) {

	client := transports.GetPassportCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		telenotification.NotifyServiceError(m.sid, m.host, m.port, nil)
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*Passport.TPassportServiceClient).GetData(context.Background(), Passport.TKey(key))

	if err != nil {
		telenotification.NotifyServiceError(m.sid, m.host, m.port, err)
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r.Data == nil {
		return nil, errors.New("Backend service:" + m.sid + " key not found")
	}
	if r.ErrorCode != Passport.TErrorCode_EGood {
		return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
	}
	return r.Data, nil
}

func (m *ppassportservice) PutData(key int64, data *Passport.TPassportInfo) error {

	client := transports.GetPassportCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*Passport.TPassportServiceClient).PutData(context.Background(), Passport.TKey(key), data)

	if err != nil {
		return errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()

	if r != Passport.TErrorCode_EGood {
		return errors.New("Backend service:" + m.sid + " err:" + r.String())
	}
	return nil
}

// func (m *userinforservice) GetMultiData(keys []int64) (map[userinfoservice.TUID]*userinfoservice.TUserInfo, error) {
// 	if m.etcdManager != nil {
// 		h, p, err := m.etcdManager.GetEndpoint(m.sid)
// 		if err != nil {
// 			log.Println("EtcdManager get endpoints", "err", err)
// 		} else {
// 			m.host = h
// 			m.port = p
// 		}
// 	}
// 	client := transports.GetUserInfoServiceCompactClient(m.host, m.port)
// 	if client == nil || client.Client == nil {
// 		return nil, errors.New("Backend service " + m.sid + "connection refused")
// 	}

// 	r, err := client.Client.(*userinfoservice.TUserInfoServiceClient).GetMultiData(context.Background(), keys)

// 	if err != nil || r == nil {
// 		return nil, errors.New("Backend service " + m.sid + "connection refused")
// 	}

// 	defer client.BackToPool()
// 	return r, nil
// 	// if r.Data == nil {
// 	// 	return nil, errors.New("Backend service:" + m.sid + " key not found")
// 	// }
// 	// if r.ErrorCode != userinfoservice.TErrorCode_EGood {
// 	// 	return nil, errors.New("Backend service:" + m.sid + " err:" + r.ErrorCode.String())
// 	// }
// 	// return r.Data, nil
// }
// func (m *userinforservice) PutData(uid int64, data *userinfoservice.TUserInfo) error {
// 	if m.etcdManager != nil {
// 		h, p, err := m.etcdManager.GetEndpoint(m.sid)
// 		if err != nil {
// 			log.Println("EtcdManager get endpoints", "err", err)
// 		} else {
// 			m.host = h
// 			m.port = p
// 		}
// 	}
// 	client := transports.GetUserInfoServiceCompactClient(m.host, m.port)
// 	if client == nil || client.Client == nil {
// 		return errors.New("Backend service " + m.sid + "connection refused")
// 	}

// 	r, err := client.Client.(*userinfoservice.TUserInfoServiceClient).PutData(context.Background(), userinfoservice.TUID(uid), data)

// 	if err != nil {
// 		return errors.New("Backend service " + m.sid + "connection refused")
// 	}

// 	defer client.BackToPool()
// 	// if r.Data == nil {
// 	// 	return nil, errors.New("Backend service:" + m.sid + " key not found")
// 	// }
// 	if r != userinfoservice.TErrorCode_EGood {
// 		return errors.New("Backend service:" + m.sid + " err:" + r.String())
// 	}
// 	return nil
// }
