package dashboard

import (
	"context"
	"github.com/B1gDaddyKane/golangBackend/src/models"
)

type ServiceDAOI interface {
	GetServices(ctx context.Context) (resp models.GetServicesResponse, err error)
	CreateService(ctx context.Context, req models.ServiceRequest) error
	UpdateService(ctx context.Context, req models.ServiceRequest) error
	DeleteService(ctx context.Context, req models.ServiceRequest) error
}
