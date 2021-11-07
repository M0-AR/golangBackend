package dao

import (
	"context"
	"errors"
	"fmt"
	"github.com/B1gDaddyKane/golangBackend/src/interfaces"
	"github.com/B1gDaddyKane/golangBackend/src/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	//"time"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceDAO struct {
	mongoDB *mongo.Database
}

func NewServiceDAO(mongo *mongo.Database) interfaces.ServiceDAOI {
	return &ServiceDAO{
		mongoDB: mongo,
	}
}

func (sdao ServiceDAO) GetServices(ctx context.Context) (resp model.GetServicesResponse, err error) {

	query, err := sdao.mongoDB.Collection("services").Find(ctx, bson.D{})
	if err != nil {
		log.Println("error", err)
		return model.GetServicesResponse{}, err
	}
	defer query.Close(ctx)

	listServices := make([]model.Service, 0)
	for query.Next(ctx) {
		var row model.Service
		err := query.Decode(&row)
		if err != nil {
			log.Println("error")
		}
		listServices = append(listServices, row)
	}

	resp = model.GetServicesResponse{Services: listServices}

	return resp, err

}

func (sdao ServiceDAO) CreateService(ctx context.Context, req model.Service) error {

	dataReq := bson.M{
		"title": req.Title,
	}

	query, err := sdao.mongoDB.Collection("services").InsertOne(ctx, dataReq)
	if err != nil {
		log.Println("error")
	}

	if oid, ok := query.InsertedID.(primitive.ObjectID); ok {
		serviceID := oid.Hex()
		updateServiceID := bson.M{"_id": oid}
		objectID := bson.M{"$set": bson.M{"_id": serviceID}}
		_, err = sdao.mongoDB.Collection("services").UpdateOne(ctx, updateServiceID, objectID)
		if err != nil {
			log.Println("error")
		}
	} else {
		err = errors.New(fmt.Sprint("can't get inserted ID", err))
		log.Println("error")
	}

	return err

}

func (sdao ServiceDAO) UpdateService(ctx context.Context, req model.Service) error {
	panic("implement me")
}

func (sdao ServiceDAO) DeleteService(ctx context.Context, req model.Service) error {
	panic("implement me")
}
