package dashboard

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/B1gDaddyKane/golangBackend/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	//"time"
	"go.mongodb.org/mongo-driver/mongo"
)

type ServiceDAO struct {
	mongoDB *mongo.Database
}

func NewServiceDAO(mongo *mongo.Database) ServiceDAOI {
	return &ServiceDAO{
		mongoDB: mongo,
	}
}

func (sdao ServiceDAO) GetServices(ctx context.Context) (resp models.GetServicesResponse, err error) {

	query, err := sdao.mongoDB.Collection("services").Find(ctx, bson.D{})
	if err != nil {
		log.Println("error ", err)
		return models.GetServicesResponse{}, err
	}
	defer query.Close(ctx)

	listServices := make([]models.Service, 0)
	for query.Next(ctx) {
		var row models.Service
		err := query.Decode(&row)
		if err != nil {
			log.Println("error ", err)
		}
		listServices = append(listServices, row)
	}

	resp = models.GetServicesResponse{Services: listServices}

	return resp, err

}

func (sdao ServiceDAO) CreateService(ctx context.Context, req models.ServiceRequest) error {

	dataReq := bson.M{
		"service_title":        req.ServiceTitle,
		"service_price":        req.ServicePrice,
		"service_image_url":    req.ServiceImageUrl,
		"service_is_available": req.ServiceIsAvailable,
		"service_start_date":   req.ServiceStartDate,
	}

	query, err := sdao.mongoDB.Collection("services").InsertOne(ctx, dataReq)
	if err != nil {
		log.Println("error")
	}

	if oid, ok := query.InsertedID.(primitive.ObjectID); ok {
		serviceID := oid.Hex()
		updateServiceID := bson.M{"_id": oid}
		objectID := bson.M{"$set": bson.M{"service_id": serviceID}}
		_, err = sdao.mongoDB.Collection("services").UpdateOne(ctx, updateServiceID, objectID)
		if err != nil {
			log.Println("error from ServiceDAO.CreateUser ", err)
		}
	} else {
		err = errors.New(fmt.Sprint("can't get inserted ID", err))
		log.Println("error")
	}

	return err

}

func (sdao ServiceDAO) UpdateService(ctx context.Context, req models.ServiceRequest) error {

	updateServiceID := bson.M{"service_id": req.ServiceID}
	objectID := bson.M{"$set": bson.M{
		"service_title": req.ServiceTitle,
	}}
	_, err := sdao.mongoDB.Collection("services").UpdateOne(ctx, updateServiceID, objectID)
	if err != nil {
		log.Println("error ", err)
	}

	return err

}

func (sdao ServiceDAO) DeleteService(ctx context.Context, req models.ServiceRequest) error {

	deleteServiceID := bson.M{"service_id": req.ServiceID}
	_, err := sdao.mongoDB.Collection("service").DeleteOne(ctx, deleteServiceID)
	if err != nil {
		log.Println("error")

	}
	return err
}
