package controller

import "REST/internal/entities"

type UseCase interface {
	GetWeather() (entities.WeatherNow, error)
}

type Controller struct {
	uc UseCase
}

func New(uc UseCase) *Controller {
	return &Controller{uc: uc}
}
