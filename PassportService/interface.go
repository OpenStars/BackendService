package PassportService

import "github.com/OpenStars/BackendService/PassportService/ppassport/thrift/gen-go/OpenStars/Platform/Passport"

type PassportService interface {
	GetData(key int64) (*Passport.TPassportInfo, error)
	PutData(key int64, data *Passport.TPassportInfo) error
}
