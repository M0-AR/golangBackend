package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Client *mongo.Client

type MongoDB struct {
	uri string
}

func NewMongoDB(uri string) *MongoDB {
	return &MongoDB{uri}
}

func (db *MongoDB) CreateUser(user UserDB) (*UserDB, error) {
	collection := Client.Database("tester").Collection("admin")
	// creates a timeout context. We give it our context and the timeout amount
	// _ means we dont handle the cancel callback.
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	x := collection.FindOne(ctx, bson.M{"_id": result})
	err = x.Decode(&user)
	return &user, err
}
