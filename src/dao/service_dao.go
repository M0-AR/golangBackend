package dao

import (
	"fmt"
	"github.com/B1gDaddyKane/golangBackend/src/errors"
	"github.com/B1gDaddyKane/golangBackend/src/model"
	"net/http"
	"time"
)

// Fake Data
var (
	services = []*model.Service{
		{ID: 1, Title: "Wash Machine", Price: 450, ImageUrl: "", IsAvailable: true, ServiceStartDate: time.Time{}},
		{ID: 2, Title: "Wash Machine02", Price: 450, ImageUrl: "", IsAvailable: true, ServiceStartDate: time.Time{}},
	}
)

func GetServices() ([]*model.Service, *errors.AppError) {
	if services != nil {
		return services, nil
	}

	return nil, &errors.AppError{
		Message:    fmt.Sprintf("Services were not found"),
		StatusCode: http.StatusNotFound,
		Status:     "not_found",
	}
}

func GetServicesById(serviceId int64) (*model.Service, error) {
	for _, item := range services {
		if item.ID == serviceId {
			return item, nil
		}
	}
	return nil, fmt.Errorf("service id not found")
}
