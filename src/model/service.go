package model

import "time"

type (
	GetServicesResponse struct {
		Services []Service `json:"services"`
	}

	Service struct {
		ServiceID          string    `json:"service_id,omitempty" bson:"service_id"`
		ServiceTitle       string    `json:"service_title" bson:"service_title"`
		ServicePrice       string    `json:"service_price" bson:"service_price"`
		ServiceImageUrl    string    `json:"service_image_url" bson:"service_image_url"`
		ServiceIsAvailable bool      `json:"service_is_available" bson:"service_is_available"`
		ServiceStartDate   time.Time `json:"service_start_date" bson:"service_start_date"`
	}

	ServiceRequest struct {
		ServiceID          string    `json:"service_id,omitempty" bson:"service_id"`
		ServiceTitle       string    `json:"service_title,omitempty" bson:"service_title"`
		ServicePrice       string    `json:"service_price,omitempty" bson:"service_price"`
		ServiceImageUrl    string    `json:"service_image_url,omitempty" bson:"service_image_url"`
		ServiceIsAvailable bool      `json:"service_is_available,omitempty" bson:"service_is_available"`
		ServiceStartDate   time.Time `json:"service_start_date,omitempty" bson:"service_start_date"`
	}
)
