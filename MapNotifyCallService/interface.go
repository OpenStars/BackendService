package MapNotifyCallService

type Client interface {
	PutData(pubkey string, token string) error
	GetTokenByPubkey(pubkey string) (string, error)
	GetPubkeyByToken(token string) (string, error)
}
