package usecase

import (
	"context"
	"github.com/B1gDaddyKane/golangBackend/lib/logging"
	"github.com/B1gDaddyKane/golangBackend/src/interfaces"
	"github.com/B1gDaddyKane/golangBackend/src/model"
	"github.com/pkg/errors"
	"log"
)

type ServiceUseCase struct {
	config     *model.EnvConfig
	serviceDAO interfaces.ServiceDAOI
}

func NewServiceUseCase(config *model.EnvConfig, serviceDAO interfaces.ServiceDAOI) interfaces.ServiceUseCaseI {
	return &ServiceUseCase{
		config:     config,
		serviceDAO: serviceDAO,
	}
}

func (suc ServiceUseCase) GetServices(ctx context.Context) (resp model.GetServicesResponse, err error) {

	if ctx == nil {
		ctx = context.Background()
	}

	list, err := suc.serviceDAO.GetServices(ctx)
	if err != nil {
		log.Println("failed to show services with default log")
		return list, err
	}

	return list, err

}

func (suc ServiceUseCase) CreateService(ctx context.Context, req model.ServiceRequest) (resp bool, err error) {

	if ctx == nil {
		ctx = context.Background()
	}

	if req.ServiceTitle == "" {
		err := errors.New("failed to add service, please add service title")
		logging.Info(err)
		return false, err
	}

	err = suc.serviceDAO.CreateService(ctx, req)
	if err != nil {
		return false, err
	}

	return true, nil

}

func (suc ServiceUseCase) UpdateService(ctx context.Context, req model.ServiceRequest) (resp bool, err error) {
	panic("implement me")
}

func (suc ServiceUseCase) DeleteService(ctx context.Context, req model.ServiceRequest) (resp bool, err error) {
	panic("implement me")
}
