package controller

import "github.com/gorilla/mux"

func (co *Controller) InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/get_weather", co.HandlerWeather).Methods("GET")
	return router
}
