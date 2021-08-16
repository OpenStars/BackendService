package MoneyStorageService

import (
	"context"
	"errors"

	"github.com/OpenStars/BackendService/MoneyStorageService/money/moneyservice"
	"github.com/OpenStars/BackendService/MoneyStorageService/money/mshared"
	"github.com/OpenStars/BackendService/MoneyStorageService/money/transports"
	telenotification "github.com/OpenStars/BackendService/TeleNotification"
)

type client struct {
	host string
	port string
	sid  string
}

func (m *client) AddMoney(uid int64, amount int64, purseType int64, uType int32, description string) (int64, error) {

	client := transports.GetMoneyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		telenotification.NotifyServiceError(m.sid, m.host, m.port, nil)
		return -1, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*moneyservice.TMoneyAgentServiceClient).AddMoney(context.Background(), uid, mshared.TMONEY(amount), mshared.TPurseType(purseType), uType, description)

	if err != nil {
		telenotification.NotifyServiceError(m.sid, m.host, m.port, nil)
		return -1, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return int64(r), nil
}

func (m *client) AddMoneyExt(uid int64, moneyType string, amount int64, purseType int64, uType int32, description string) (int64, error) {

	client := transports.GetMoneyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return -1, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*moneyservice.TMoneyAgentServiceClient).AddMoneyExt(context.Background(), uid, moneyType, mshared.TMONEY(amount), mshared.TPurseType(purseType), uType, description)

	if err != nil {
		return -1, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return int64(r), nil
}

func (m *client) CanTakeMoney(uid int64, uType int32, purseType int64, amount int64) (*mshared.TAccountPurses, error) {

	client := transports.GetMoneyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*moneyservice.TMoneyAgentServiceClient).CanTakeMoney(context.Background(), uid, uType, mshared.TPurseType(purseType), mshared.TMONEY(amount))

	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	if r == nil {
		return nil, nil
	}
	return r.Purses, nil
}

func (m *client) CreateMoneyAccountFor(uid int64, uType int32, creditLimit, hardCreditLimit int64) (bool, error) {

	client := transports.GetMoneyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return false, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*moneyservice.TMoneyAgentServiceClient).CreateMoneyAccountFor(context.Background(), uid, uType, mshared.TMONEY(creditLimit), mshared.TMONEY(hardCreditLimit))

	if err != nil {
		return false, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return r, nil
}

func (m *client) GetMoney(uid int64, uType int32) (*moneyservice.TPurses, error) {

	client := transports.GetMoneyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*moneyservice.TMoneyAgentServiceClient).GetMoney(context.Background(), uid, uType)

	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return r, nil
}
func (m *client) GetTransactionsByCount(uid int64, uType, pos, count int32) ([]*moneyservice.TTransaction, error) {

	client := transports.GetMoneyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*moneyservice.TMoneyAgentServiceClient).GetTransactionsByCount(context.Background(), uid, uType, pos, count)

	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return r, nil
}

func (m *client) GetTransactionsByCountR(uid int64, uType, pos, count int32) ([]*moneyservice.TTransaction, error) {

	client := transports.GetMoneyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*moneyservice.TMoneyAgentServiceClient).GetTransactionsByCountR(context.Background(), uid, uType, pos, count)

	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return r, nil
}
func (m *client) GetTransactionsByRange(uid int64, uType int32, fromTime, toTime int64) ([]*moneyservice.TTransaction, error) {

	client := transports.GetMoneyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*moneyservice.TMoneyAgentServiceClient).GetTransactionsByRange(context.Background(), uid, uType, mshared.TTIME(fromTime), mshared.TTIME(toTime))

	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return r, nil
}

func (m *client) GetTransactionsBySlice(uid int64, uType int32, fromTime int64, count int32) ([]*moneyservice.TTransaction, error) {

	client := transports.GetMoneyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*moneyservice.TMoneyAgentServiceClient).GetTransactionsBySlice(context.Background(), uid, uType, mshared.TTIME(fromTime), count)

	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return r, nil
}

func (m *client) TransferMoney(fromUser int64, fromPurse int64, fromUType int32, toUser int64, toPurse int64, toUType int32, amount int64, description string) (int64, error) {

	client := transports.GetMoneyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return -1, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*moneyservice.TMoneyAgentServiceClient).TransferMoney(context.Background(), fromUser, mshared.TPurseType(fromPurse), fromUType, toUser, mshared.TPurseType(toPurse), toUType, mshared.TMONEY(amount), description)

	if err != nil {
		return -1, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return int64(r), nil
}

func (m *client) TransferMoneyExt(fromUser int64, fromPurse int64, fromUType int32, fromMoneyTypeExt string, toUser int64, toPurse int64, toUType int32, toMoneyTypeExt string, amount int64, description string) (int64, error) {

	client := transports.GetMoneyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return -1, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*moneyservice.TMoneyAgentServiceClient).TransferMoneyExt(context.Background(), fromUser, mshared.TPurseType(fromPurse), fromUType, fromMoneyTypeExt, toUser, mshared.TPurseType(toPurse), toUType, toMoneyTypeExt, mshared.TMONEY(amount), description)

	if err != nil {
		return -1, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return int64(r), nil
}

func (m *client) GetTransactionResult_(transactionId int64) (*mshared.TTransactionResult_, error) {
	client := transports.GetMoneyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*moneyservice.TMoneyAgentServiceClient).GetTransactionResult_(context.Background(), mshared.TTRANSID(transactionId))

	if err != nil {
		return nil, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return r, nil
}
func (m *client) UpdateCreditLimit(uID int64, uType int32, creditLimit int64, hardCreditLimit int64) (bool, error) {
	client := transports.GetMoneyStorageServiceCompactClient(m.host, m.port)
	if client == nil || client.Client == nil {
		return false, errors.New("Backend service " + m.sid + "connection refused")
	}

	r, err := client.Client.(*moneyservice.TMoneyAgentServiceClient).UpdateCreditLimit(context.Background(), uID, uType, mshared.TMONEY(creditLimit), mshared.TMONEY(hardCreditLimit))

	if err != nil {
		return false, errors.New("Backend service " + m.sid + "connection refused")
	}

	defer client.BackToPool()
	return r, nil
}
