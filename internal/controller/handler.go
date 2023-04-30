package controller

import (
	"encoding/json"
	"net/http"
)

func (co *Controller) HandlerWeather(w http.ResponseWriter, r *http.Request) {
	weatherNow, err := co.uc.GetWeather()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsResponse, err := json.Marshal(weatherNow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsResponse)
}
