package usecases

import "time"

type Thermometer interface {
	GetTemperature() (float32, error)
}
type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, duration time.Duration)
}
type UseCase struct {
	tm    Thermometer
	cache Cache
}

func New(tm Thermometer, cache Cache) *UseCase {
	return &UseCase{tm: tm, cache: cache}
}
