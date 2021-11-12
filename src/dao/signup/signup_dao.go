package signup

import (
	"context"
	"github.com/B1gDaddyKane/golangBackend/src/models"
	"github.com/B1gDaddyKane/golangBackend/src/security"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var Client *mongo.Client

type SignUpDAO struct {
	mongoDB *mongo.Database
}

func NewSignUpDAO(mongo *mongo.Database) SignUpDAOI {
	return &SignUpDAO{
		mongoDB: mongo,
	}
}

func (su SignUpDAO) GetUser(ctx context.Context) (resp models.GetUserResponse, err error) {
	panic("implement me")
}

func (su SignUpDAO) CreateUser(ctx context.Context, req models.UserRequest) error {

	dataReq := bson.M{
		"Firstname": req.Firstname,
		"Lastname":  req.Lastname,
		"Email":     req.Email,
		"Password":  security.GetHash([]byte(req.Password)),
		"Created":   time.Now().Unix(),
	}

	collection := su.mongoDB.Collection("users")
	// creates a timeout context. We give it our context and the timeout amount
	// _ means we dont handle the cancel callback.
	result, err := collection.InsertOne(ctx, dataReq)
	if err != nil {
		return err
	}
	x := collection.FindOne(ctx, bson.M{"_id": result})
	err = x.Decode(&req)
	return err
}

func (su SignUpDAO) UpdateUser(ctx context.Context, req models.UserRequest) error {
	panic("implement me")
}

func (su SignUpDAO) DeleteUser(ctx context.Context, req models.UserRequest) error {
	panic("implement me")
}
