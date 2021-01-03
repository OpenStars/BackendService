package MoneyStorageService

import (
	"github.com/OpenStars/BackendService/MoneyStorageService/money/moneyservice"
	"github.com/OpenStars/BackendService/MoneyStorageService/money/mshared"
)

type Client interface {
	// Parameters:
	//  - UID
	//  - UType
	//  - CreditLimit
	//  - HardCreditLimit
	CreateMoneyAccountFor(uID int64, uType int32, creditLimit int64, hardCreditLimit int64) (r bool, err error)
	// Parameters:
	//  - UID
	//  - UType
	//  - CreditLimit
	//  - HardCreditLimit
	UpdateCreditLimit(uID int64, uType int32, creditLimit int64, hardCreditLimit int64) (r bool, err error)
	// Parameters:
	//  - UID
	//  - UType
	//  - PurseType
	//  - Amount
	CanTakeMoney(uID int64, uType int32, purseType int64, amount int64) (r *mshared.TAccountPurses, err error)
	// Parameters:
	//  - UID
	//  - UType
	GetMoney(uID int64, uType int32) (r *moneyservice.TPurses, err error)
	// Parameters:
	//  - UID
	//  - Amount
	//  - PurseType
	//  - UType
	//  - Description
	AddMoney(uID int64, amount int64, purseType int64, uType int32, description string) (r int64, err error)
	// Parameters:
	//  - UID
	//  - MoneyType
	//  - Amount
	//  - PurseType
	//  - UType
	//  - Description
	AddMoneyExt(uID int64, moneyType string, amount int64, purseType int64, uType int32, description string) (r int64, err error)
	// Parameters:
	//  - FromUser
	//  - FromPurse
	//  - FromUType
	//  - ToUser
	//  - ToPurse
	//  - ToUType
	//  - Amount
	//  - Description
	TransferMoney(fromUser int64, fromPurse int64, fromUType int32, toUser int64, toPurse int64, toUType int32, amount int64, description string) (r int64, err error)
	// Parameters:
	//  - FromUser
	//  - FromPurse
	//  - FromUType
	//  - FromMoneyTypeExt
	//  - ToUser
	//  - ToPurse
	//  - ToUType
	//  - ToMoneyTypeExt
	//  - Amount
	//  - Description
	TransferMoneyExt(fromUser int64, fromPurse int64, fromUType int32, fromMoneyTypeExt string, toUser int64, toPurse int64, toUType int32, toMoneyTypeExt string, amount int64, description string) (r int64, err error)
	// Parameters:
	//  - UID
	//  - UType
	//  - FromTime
	//  - ToTime
	GetTransactionsByRange(uid int64, uType int32, fromTime, toTime int64) (r []*moneyservice.TTransaction, err error)
	// Parameters:
	//  - UID
	//  - UType
	//  - FromTime
	//  - Count
	GetTransactionsBySlice(uid int64, uType int32, fromTime int64, count int32) (r []*moneyservice.TTransaction, err error)
	// Parameters:
	//  - UID
	//  - UType
	//  - Pos
	//  - Count
	GetTransactionsByCount(uID int64, uType int32, pos int32, count int32) (r []*moneyservice.TTransaction, err error)
	// Parameters:
	//  - UID
	//  - UType
	//  - Pos
	//  - Count
	GetTransactionsByCountR(uID int64, uType int32, pos int32, count int32) (r []*moneyservice.TTransaction, err error)
	// Parameters:
	//  - TransactionId
	GetTransactionResult_(transactionId int64) (r *mshared.TTransactionResult_, err error)
}
