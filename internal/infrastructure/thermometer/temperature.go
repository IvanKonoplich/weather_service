package thermometer

import (
	"math/rand"
	"time"
)

func (tm *Thermometer) GetTemperature() (float32, error) {
	//Имитируем возврат какой то температуры
	rand.Seed(time.Now().UnixNano())
	n := rand.Float32()
	return n, nil
}
