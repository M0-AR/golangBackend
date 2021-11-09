package dashboard

import (
	"context"
	"github.com/B1gDaddyKane/golangBackend/lib/logging"
	"github.com/B1gDaddyKane/golangBackend/src/dao/dashboard"
	"github.com/B1gDaddyKane/golangBackend/src/models"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"log"
)

type ServiceUseCase struct {
	config     *models.EnvConfig
	serviceDAO dashboard.ServiceDAOI
}

func NewServiceUseCase(config *models.EnvConfig, serviceDAO dashboard.ServiceDAOI) ServiceUseCaseI {
	return &ServiceUseCase{
		config:     config,
		serviceDAO: serviceDAO,
	}
}

func (suc ServiceUseCase) GetServices(ctx context.Context) (resp models.GetServicesResponse, err error) {

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

func (suc ServiceUseCase) CreateService(ctx context.Context, req models.ServiceRequest) (resp bool, err error) {

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

func (suc ServiceUseCase) UpdateService(ctx context.Context, req models.ServiceRequest) (resp bool, err error) {

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

func (suc ServiceUseCase) DeleteService(ctx context.Context, req models.ServiceRequest) (resp bool, err error) {
	panic("implement me")
}
