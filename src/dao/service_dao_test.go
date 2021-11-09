package dao

//
//import (
//	"github.com/stretchr/testify/assert"
//	"net/http"
//	"strconv"
//	"testing"
//)
//
///** TestGetServices */
//func TestGetServicesOnFailScenario(t *testing.T) {
//
//	services, err := GetServices()
//
//	assert.Nil(t, services)
//	assert.NotNil(t, err)
//	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
//	assert.EqualValues(t, "not_found", err.Status)
//	assert.EqualValues(t, "Services were not found", err.Message)
//
//}
//
//func TestGetServicesOnSuccessScenario(t *testing.T) {
//
//	servicesLength := 2
//
//	services, err := GetServices()
//
//	assert.Nil(t, err)
//	assert.NotNil(t, services)
//	assert.EqualValues(t, servicesLength, len(services))
//
//}
//
///** TestGetServiceById */
//
//func TestGetServiceByIdOnFailScenario(t *testing.T) {
//
//	var serviceId int64 = 321321
//	service, err := GetServiceById(serviceId)
//
//	assert.Nil(t, service)
//	assert.NotNil(t, err)
//	assert.EqualValues(t, http.StatusNotFound, err.StatusCode)
//	assert.EqualValues(t, "not_found", err.Status)
//	assert.EqualValues(t, "Service "+strconv.Itoa(int(serviceId))+" was not found", err.Message)
//
//}
//
//func TestGetServiceByIdOnSuccessScenario(t *testing.T) {
//
//	var httpCode = 200
//
//	var serviceId int64 = 1
//	var serviceTitle = "Wash Machine"
//
//	service, err := GetServiceById(serviceId)
//
//	assert.Nil(t, err)
//	assert.NotNil(t, service)
//	assert.EqualValues(t, httpCode, http.StatusOK)
//	assert.EqualValues(t, serviceId, service.ID)
//	assert.EqualValues(t, serviceTitle, service.Title)
//
//}
