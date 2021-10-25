package controllers

import (
	"encoding/json"
	"github.com/B1gDaddyKane/golangBackend/src/services"
	"net/http"
)

// GetServices Get all services
func GetServices(response http.ResponseWriter, request *http.Request) {
	services, apiErr := services.GetServices()

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(services)
	response.Write(jsonValue)
}

func GetServiceById(response http.ResponseWriter, request *http.Request) {

}
