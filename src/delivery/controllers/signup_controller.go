package controllers

import (
	"context"
	"encoding/json"
	"github.com/B1gDaddyKane/golangBackend/src/models"
	"github.com/B1gDaddyKane/golangBackend/src/usecase/signup"
	"github.com/labstack/echo/v4"
)

type SignUpController struct {
	e        *echo.Echo
	sUseCase signup.SignUpUseCaseI
}

func NewSignUpController(e *echo.Echo, sUseCase signup.SignUpUseCaseI) *SignUpController {
	return &SignUpController{
		e:        e,
		sUseCase: sUseCase,
	}
}

func (suc *SignUpController) CreateUser(ec echo.Context) error {

	var req models.UserRequest
	err := json.NewDecoder(ec.Request().Body).Decode(&req)
	if err != nil {
		return err
	}
	data, err := suc.sUseCase.CreateUser(context.Background(), req)
	if err != nil {
		return err
	}

	return ec.JSON(200, data) // Todo: add status code instead of 200

}
