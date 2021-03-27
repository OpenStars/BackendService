package KVCounterService

import "github.com/OpenStars/BackendService/KVCounterService/kvcounter/thrift/gen-go/OpenStars/Counters/KVStepCounter"

type Client interface {
	GetValue(genname string) (int64, error)
	CreateGenerator(genname string) (int32, error)
	GetStepValue(genname string, step int64) (int64, error)
	GetCurrentValue(genname string) (int64, error)
	GetMultiValue(listKeys []string) ([]*KVStepCounter.TKVCounterItem, error)
	GetMultiCurrentValue(listKeys []string) ([]*KVStepCounter.TKVCounterItem, error)
	Close()
}
