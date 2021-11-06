package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/B1gDaddyKane/golangBackend/src/errors"
	"github.com/B1gDaddyKane/golangBackend/src/services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// GetServices Get all services
func GetServices(response http.ResponseWriter, request *http.Request) {
	fmt.Print("GetServices is being invoked\n")

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
	fmt.Print("GetServiceById is being invoked\n")

	params := mux.Vars(request)

	serviceId, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {

		apiErr := &errors.AppError{
			Message:    fmt.Sprintf("id must be a number"),
			StatusCode: http.StatusBadRequest,
			Status:     "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}

	service, apiErr := services.GetServiceById(serviceId)

	if apiErr != nil {

		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)
		return
	}

	jsonValue, _ := json.Marshal(service)
	response.Write(jsonValue)
}
