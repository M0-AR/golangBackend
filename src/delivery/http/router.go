package http

import (
	"github.com/B1gDaddyKane/golangBackend/src/delivery/controllers"
	"github.com/B1gDaddyKane/golangBackend/src/interfaces"
	"github.com/labstack/echo/v4"
)

func NewRouter(e *echo.Echo, serviceUseCase interfaces.ServiceUseCaseI) {

	// Dashboard -> Services
	serviceCtrl := controllers.NewServiceController(e, serviceUseCase)

	r := e.Group("/api/v1/go-mongo")
	r.GET("/dashboard/list", serviceCtrl.GetServices)
	r.POST("/dashboard/add", serviceCtrl.CreateService)
	r.PUT("/dashboard/edit", serviceCtrl.UpdateService)

}
