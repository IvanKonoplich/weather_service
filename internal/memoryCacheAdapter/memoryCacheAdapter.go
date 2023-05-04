package memoryCacheAdapter

import (
	"time"
)

type MemoryCache interface {
	Set(key string, value interface{}, duration time.Duration)
	Get(key string) (interface{}, bool)
	Delete(key string) error
}

type MemoryCacheAdapter struct {
	mc MemoryCache
}

func New(mc MemoryCache) *MemoryCacheAdapter {
	return &MemoryCacheAdapter{
		mc: mc,
	}
}

func (mca MemoryCacheAdapter) Get() (float32, bool, error) {
	ct, ok := mca.mc.Get("temperature")
	if !ok {
		return 0.0, ok, nil
	}
	ctF := ct.(float32)
	return ctF, ok, nil
}
func (mca MemoryCacheAdapter) Set(temperature float32) error {
	mca.mc.Set("temperature", temperature, 10*time.Second)
	return nil
}
