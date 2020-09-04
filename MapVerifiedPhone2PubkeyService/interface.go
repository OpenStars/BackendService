package MapVerifiedPhone2PubkeyService

type Client interface {
	PutData(pubkey string, phonenumber string) (bool, error)
	GetPhoneNumberByPubkey(pubkey string) (string, error)
	GetPubkeyByPhoneNumber(phonenumber string) (string, error)
}
