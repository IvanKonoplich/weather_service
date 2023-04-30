package app

import (
	controller2 "REST/internal/controller"
	thermometer2 "REST/internal/infrastructure/thermometer"
	"REST/internal/usecases"
	"github.com/sirupsen/logrus"
)

func RunApp() {
	thermometer := thermometer2.New()
	cache:=
	useCase := usecases.New(thermometer)
	controller := controller2.New(useCase)
	server := new(controller2.Server)
	router := controller.InitRouter()
	err := server.RunServer("8080", router)
	logrus.Fatalf("error while starting server:%s", err.Error())

	logrus.Info("starting server...")
}
