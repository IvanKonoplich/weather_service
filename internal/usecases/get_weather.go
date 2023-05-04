package usecases

import "REST/internal/entities"

func (uc *UseCase) GetWeather() (entities.WeatherNow, error) {
	currentTemperature, ok, err := uc.cache.Get()
	if err != nil {
		return entities.WeatherNow{}, err
	}
	if ok {
		return entities.WeatherNow{Temperature: currentTemperature * 100}, nil
	}
	currentTemperature, err = uc.tm.GetTemperature()
	if err != nil {
		return entities.WeatherNow{}, err
	}
	err = uc.cache.Set(currentTemperature)
	if err != nil {
		return entities.WeatherNow{}, err
	}
	return entities.WeatherNow{Temperature: currentTemperature * 100}, nil
}
