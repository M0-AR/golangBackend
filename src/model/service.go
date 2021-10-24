package model

import (
	"time"
)

// Service represents data about a record Service.
type Service struct {
	ID               int       `json:"id"`
	Title            string    `json:"title"`
	Price            float64   `json:"price"`
	ImageUrl         string    `json:"image_url"`
	IsAvailable      bool      `json:"is_available"`
	ServiceStartDate time.Time `json:"service_start_date"`
}
