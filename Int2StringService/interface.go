package Int2StringService

type Client interface {
	PutData(key int64, value string) (bool, error)
	GetData(key int64) (string, error)
}
