package UserInfoService

import "github.com/OpenStars/BackendService/UserInfoService/userinfoservice/thrift/gen-go/openstars/userinfoservice"

type UserInfoService interface {
	GetData(key int64) (*userinfoservice.TUserInfo, error)
	PutData(uid int64, data *userinfoservice.TUserInfo) error
	GetMultiData(keys []int64) (map[userinfoservice.TUID]*userinfoservice.TUserInfo, error)
}
