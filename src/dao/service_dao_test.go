package dao

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strconv"
	"testing"
)

func TestGetServicesByIdOnFailScenario(t *testing.T) {

	var serviceId int64 = 321321
	service, err := GetServicesById(serviceId)

	assert.Nil(t, service)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
	assert.EqualValues(t, "not_found", err.Status)
	assert.EqualValues(t, "Service "+strconv.Itoa(int(serviceId))+" was not found", err.Message)

}
