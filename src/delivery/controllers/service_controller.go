package controllers

import (
	"context"
	"encoding/json"

	"github.com/B1gDaddyKane/golangBackend/src/models"
	"github.com/B1gDaddyKane/golangBackend/src/usecase/dashboard"
	"github.com/labstack/echo/v4"
)

type ServiceController struct {
	e        *echo.Echo
	sUseCase dashboard.ServiceUseCaseI
}

func NewServiceController(e *echo.Echo, sUseCase dashboard.ServiceUseCaseI) *ServiceController {
	return &ServiceController{
		e:        e,
		sUseCase: sUseCase,
	}
}

func (sc *ServiceController) GetServices(ec echo.Context) error {

	data, err := sc.sUseCase.GetServices(context.Background())
	if err != nil {
		return err
	}

	return ec.JSON(200, data) // Todo: add status code instead of 200

}

func (sc *ServiceController) CreateService(ec echo.Context) error {

	var req models.ServiceRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return err
	}
	data, err := sc.sUseCase.CreateService(context.Background(), req)
	if err != nil {
		return err
	}

	return ec.JSON(200, data) // Todo: add status code instead of 200

}

func (sc *ServiceController) UpdateService(ec echo.Context) error {

	var req models.ServiceRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return nil
	}
	data, err := sc.sUseCase.UpdateService(context.Background(), req)
	if err != nil {
		return err
	}

	return ec.JSON(200, data)
}
func (sc *ServiceController) DeleteService(ec echo.Context) error {

	var req models.ServiceRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return nil
	}
	data, err := sc.sUseCase.DeleteService(context.Background(), req)
	if err != nil {
		return err
	}

	return ec.JSON(200, data)

}
