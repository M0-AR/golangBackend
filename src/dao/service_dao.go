package dao

import (
	"fmt"
	"github.com/B1gDaddyKane/golangBackend/src/model"
	"time"
)

// Fake Data
var (
	services = []model.Service{
		{ID: 1, Title: "Wash Machine", Price: 450, ImageUrl: "", IsAvailable: true, ServiceStartDate: time.Time{}},
		{ID: 2, Title: "Wash Machine02", Price: 450, ImageUrl: "", IsAvailable: true, ServiceStartDate: time.Time{}},
	}
)

func GetServices() ([]model.Service, error) {
	if services != nil {
		return services, nil
	}

	return nil, fmt.Errorf("services are empty")
}

func GetServicesById(serviceId int) (*model.Service, error) {
	for _, item := range services {
		if item.ID == serviceId {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("service id not found")
}
