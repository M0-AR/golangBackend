package signup

import (
	"context"
	"github.com/B1gDaddyKane/golangBackend/src/models"
)

type SignUpDAOI interface {
	GetUser(ctx context.Context) (resp models.GetUserResponse, err error)
	CreateUser(ctx context.Context, req models.UserRequest) error
	UpdateUser(ctx context.Context, req models.UserRequest) error
	DeleteUser(ctx context.Context, req models.UserRequest) error
}
