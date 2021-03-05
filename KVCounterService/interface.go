package KVCounterService

type Client interface {
	GetValue(genname string) (int64, error)
	CreateGenerator(genname string) (int32, error)
	GetStepValue(genname string, step int64) (int64, error)
	GetCurrentValue(genname string) (int64, error)

	Close()
}
