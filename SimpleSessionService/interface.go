package SimpleSessionService

import "github.com/OpenStars/BackendService/SimpleSessionService/simplesession/thrift/gen-go/simplesession"

type Client interface {

	// ================================= V2 ===========================================
	GetSession(sskey string) (*simplesession.TUserSessionInfo, error)
	CreateSession(uid int64, deviceInfo string, data string, expiredTime int64) (sessionkey string, err error)
	RemoveSession(sskey string) (bool, error)
}
