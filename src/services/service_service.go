package services

import (
	"github.com/B1gDaddyKane/golangBackend/src/dao"
	"github.com/B1gDaddyKane/golangBackend/src/errors"
	"github.com/B1gDaddyKane/golangBackend/src/model"
)

func GetServices() ([]*model.Service, *errors.AppError) {
	return dao.GetServices()
}

func GetServiceById(serviceId int64) (*model.Service, *errors.AppError) {
	return dao.GetServiceById(serviceId)
}
