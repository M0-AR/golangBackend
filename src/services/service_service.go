package services

import (
	"github.com/B1gDaddyKane/golangBackend/src/dao"
	"github.com/B1gDaddyKane/golangBackend/src/model"
)

func GetServices() ([]model.Service, error) {
	services, err := dao.GetServices()

	if err != nil {
		return nil, err
	}

	return services, nil
}

func GetServiceById(serviceId int) (*model.Service, error) {
	service, err := dao.GetServicesById(serviceId)

	if err != nil {
		return nil, err
	}

	return service, nil
}
