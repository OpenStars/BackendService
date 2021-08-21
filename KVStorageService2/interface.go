package KVStorageService

import "github.com/OpenStars/BackendService/KVStorageService2/kvstorage/thrift/gen-go/OpenStars/Platform/KVStorage"

type Client interface {
	GetData(key string) ([]byte, error)
	PutData(key string, value []byte) (bool, error)
	RemoveData(key string) (bool, error)
	GetListData(keys []string) (results []*KVStorage.KVItem, missingkeys []string, err error)
}
