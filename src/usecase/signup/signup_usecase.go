package signup

import (
	"context"
	"github.com/B1gDaddyKane/golangBackend/src/dao/signup"
	"github.com/B1gDaddyKane/golangBackend/src/models"
)

type SignUpUseCase struct {
	config    *models.EnvConfig
	signupDAO signup.SignUpDAOI
}

func NewSignUpUseCase(config *models.EnvConfig, signupDAO signup.SignUpDAOI) SignUpUseCaseI {
	return &SignUpUseCase{
		config:    config,
		signupDAO: signupDAO,
	}
}

func (suuc SignUpUseCase) GetUser(ctx context.Context) (resp models.GetUserResponse, err error) {
	panic("implement me")
}

func (suuc SignUpUseCase) CreateUser(ctx context.Context, req models.UserRequest) (resp bool, err error) {
	if ctx == nil {
		ctx = context.Background()
	}

	//if req.ProductName == "" {
	//	err = errors.New("failed to add data product ")
	//	logging.Info(err)
	//	return false, err
	//}

	//insert data
	err = suuc.signupDAO.CreateUser(ctx, req)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (suuc SignUpUseCase) UpdateUser(ctx context.Context, req models.UserRequest) (resp bool, err error) {
	panic("implement me")
}

func (suuc SignUpUseCase) DeleteUser(ctx context.Context, req models.UserRequest) (resp bool, err error) {
	panic("implement me")
}
