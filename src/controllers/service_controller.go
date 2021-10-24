package controllers

import (
	"encoding/json"
	"github.com/B1gDaddyKane/golangBackend/src/services"
	"net/http"
)

// GetServices Get all services
func GetServices(response http.ResponseWriter, request *http.Request) {
	services, err := services.GetServices()

	if err != nil {
		response.WriteHeader(http.StatusNotFound) // Todo
		response.Write([]byte(err.Error()))
		return
	}

	jsonValue, _ := json.Marshal(services)
	response.Write(jsonValue)
}

func GetServiceById(response http.ResponseWriter, request *http.Request) {

}
