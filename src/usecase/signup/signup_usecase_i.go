package signup

import (
	"context"
	"github.com/B1gDaddyKane/golangBackend/src/models"
)

type SignUpUseCaseI interface {
	GetUser(ctx context.Context) (resp models.GetUserResponse, err error)
	CreateUser(ctx context.Context, req models.UserRequest) (resp bool, err error)
	UpdateUser(ctx context.Context, req models.UserRequest) (resp bool, err error)
	DeleteUser(ctx context.Context, req models.UserRequest) (resp bool, err error)
}
