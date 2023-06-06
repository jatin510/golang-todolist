package database

import "go.mongodb.org/mongo-driver/mongo"

func InitDatabase() (*mongo.Client, *mongo.Database) {
	return &mongo.Client{}, &mongo.Database{}
}
