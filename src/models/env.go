package models

import "github.com/B1gDaddyKane/golangBackend/lib/db"

type (
	EnvConfig struct {
		Host  string
		Port  int
		Mongo db.MongoConfig
	}
)
