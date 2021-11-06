package dao

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

/** TestGetServiceById */

func TestGetServiceByIdOnFailScenario(t *testing.T) {

	var serviceId int64 = 321321
	service, err := GetServiceById(serviceId)

	assert.Nil(t, service)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Status)
	assert.EqualValues(t, "Service "+strconv.Itoa(int(serviceId))+" was not found", err.Message)

}

func TestGetServiceByIdOnSuccessScenario(t *testing.T) {

	var httpCode = 200

	var serviceId int64 = 1
	var serviceTitle = "Wash Machine"

	service, err := GetServiceById(serviceId)

	assert.Nil(t, err)
	assert.NotNil(t, service)
	assert.EqualValues(t, httpCode, http.StatusOK)
	assert.EqualValues(t, serviceId, service.ID)
	assert.EqualValues(t, serviceTitle, service.Title)

}
