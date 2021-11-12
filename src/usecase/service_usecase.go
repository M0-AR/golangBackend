package usecase

import (
	"context"
	"github.com/B1gDaddyKane/golangBackend/lib/logging"
	"github.com/B1gDaddyKane/golangBackend/src/dao"
	"github.com/B1gDaddyKane/golangBackend/src/model"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"log"
)

type ServiceUseCase struct {
	config     *model.EnvConfig
	serviceDAO dao.ServiceDAOI
}

func NewServiceUseCase(config *model.EnvConfig, serviceDAO dao.ServiceDAOI) ServiceUseCaseI {
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

	if req.ServiceTitle == "" || len(req.ServicePrice) < 2 || req.ServiceImageUrl == "" { // Todo: should I validate characters like price should just number
		err := errors.New("failed to add service, please fill service's attributes")
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

	if ctx == nil {
		ctx = context.Background()
	}

	err = suc.serviceDAO.UpdateService(ctx, req)
	if err != nil {
		return false, err
	}

	zapLogger, _ := zap.NewProduction()
	defer zapLogger.Sync()
	zapLogger.Info("success to update service",
		zap.String("service ID", req.ServiceID),
	)

	return true, nil // Todo: Avoid returning nil

}

func (suc ServiceUseCase) DeleteService(ctx context.Context, req model.ServiceRequest) (resp bool, err error) {
	panic("implement me")
}
