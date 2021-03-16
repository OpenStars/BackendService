package KVStorageService

type Client interface {
	GetData(key string) (string, error)
	PutData(key string, value string) (bool, error)
	RemoveData(key string) (bool, error)
	GetListData(keys []string) (results map[string]string, missingkeys []string, err error)
	Close()
}
