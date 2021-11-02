package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserDB struct {
	ID        primitive.ObjectID `bson:"_id"`
	Firstname string             `bson:"firstname"`
	Lastname  string             `bson:"lastname"`
	Email     string             `bson:"email"`
	Password  string             `bson:"hashedPassword"`
	Created   int64              `bson:"created"`
}

type Database interface {
	CreateUser(user UserDB) (*UserDB, error)
	UpdateUser(user UserDB) (*UserDB, error)
	DeleteUser(user UserDB) (*UserDB, error)
	GetUser(user UserDB) (*UserDB, error)
}
