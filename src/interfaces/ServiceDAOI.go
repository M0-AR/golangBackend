package interfaces

import (
	"context"
	"github.com/B1gDaddyKane/golangBackend/src/model"
)

type ServiceDAOI interface {
	GetServices(ctx context.Context) (resp model.GetServicesResponse, err error)
	CreateService(ctx context.Context, req model.Service) error
	UpdateService(ctx context.Context, req model.Service) error
	DeleteService(ctx context.Context, req model.Service) error
}
