package usecases

type Thermometer interface {
	GetTemperature() (float32, error)
}
type Cache interface {
	Get() (float32, bool, error)
	Set(temperature float32) error
}
type UseCase struct {
	tm    Thermometer
	cache Cache
}

func New(tm Thermometer, cache Cache) *UseCase {
	return &UseCase{tm: tm, cache: cache}
}
