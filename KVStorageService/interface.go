package KVStorageService

import "github.com/OpenStars/BackendService/KVStorageService/kvstorage/thrift/gen-go/OpenStars/Platform/KVStorage"

type Client interface {
	GetData(key string) (string, error)
	PutData(key string, value string) (bool, error)
	RemoveData(key string) (bool, error)
	GetListData(keys []string) (results []*KVStorage.KVItem, missingkeys []string, err error)
	NextItem(sessionKey int64) (*KVStorage.KVItem, error)
	OpenIterate() (int64, error)
	CloseIterate(sessionKey int64) error
	NextListItems(sessionKey, count int64) ([]*KVStorage.KVItem, error)
}
