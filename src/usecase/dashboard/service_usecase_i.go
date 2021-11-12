package dashboard

import (
	"context"
	"github.com/B1gDaddyKane/golangBackend/src/models"
)

type ServiceUseCaseI interface {
	GetServices(ctx context.Context) (resp models.GetServicesResponse, err error)
	CreateService(ctx context.Context, req models.ServiceRequest) (resp bool, err error)
	UpdateService(ctx context.Context, req models.ServiceRequest) (resp bool, err error)
	DeleteService(ctx context.Context, req models.ServiceRequest) (resp bool, err error)
}
