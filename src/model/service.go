package model

import (
	"time"
)

type (
	GetServicesResponse struct {
		Services []Service `json:"services"`
	}

	Service struct {
		ID               int64     `json:"id,omitempty" bson:"_id"`
		Title            string    `json:"title" bson:"title"`
		Price            float64   `json:"price" bson:"price"`
		ImageUrl         string    `json:"image_url" bson:"image_url"`
		IsAvailable      bool      `json:"is_available" bson:"is_available"`
		ServiceStartDate time.Time `json:"service_start_date" bson:"service_start_date"`
	}

	ServiceRequest struct {
		ID               int64     `json:"id,omitempty" bson:"_id"`
		Title            string    `json:"title,omitempty" bson:"title"`
		Price            float64   `json:"price,omitempty" bson:"price"`
		ImageUrl         string    `json:"image_url,omitempty" bson:"image_url"`
		IsAvailable      bool      `json:"is_available,omitempty" bson:"is_available"`
		ServiceStartDate time.Time `json:"service_start_date,omitempty" bson:"service_start_date"`
	}
)
