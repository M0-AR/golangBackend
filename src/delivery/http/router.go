package http

import (
	"github.com/B1gDaddyKane/golangBackend/src/delivery/controllers"
	"github.com/B1gDaddyKane/golangBackend/src/usecase/dashboard"
	"github.com/B1gDaddyKane/golangBackend/src/usecase/signup"
	"github.com/labstack/echo/v4"
)

func NewRouterSignUp(e *echo.Echo, signupUseCase signup.SignUpUseCaseI) {

	// Signup -> User
	signupCtrl := controllers.NewSignUpController(e, signupUseCase)

	r := e.Group("/api/v1/go-mongo")
	r.POST("/signup", signupCtrl.CreateUser)
}

func NewRouterDashboard(e *echo.Echo, serviceUseCase dashboard.ServiceUseCaseI) {

	// dashboard -> Services
	serviceCtrl := controllers.NewServiceController(e, serviceUseCase)

	r := e.Group("/api/v1/go-mongo")
	r.GET("/dashboard/list", serviceCtrl.GetServices)
	r.POST("/dashboard/add", serviceCtrl.CreateService)
	r.PUT("/dashboard/edit", serviceCtrl.UpdateService)
	r.POST("/dashboard/delete", serviceCtrl.DeleteService)

	//

}
