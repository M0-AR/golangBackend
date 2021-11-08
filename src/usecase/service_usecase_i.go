package usecase

import (
	"context"
	"github.com/B1gDaddyKane/golangBackend/src/model"
)

type ServiceUseCaseI interface {
	GetServices(ctx context.Context) (resp model.GetServicesResponse, err error)
	CreateService(ctx context.Context, req model.ServiceRequest) (resp bool, err error)
	UpdateService(ctx context.Context, req model.ServiceRequest) (resp bool, err error)
	DeleteService(ctx context.Context, req model.ServiceRequest) (resp bool, err error)
}
