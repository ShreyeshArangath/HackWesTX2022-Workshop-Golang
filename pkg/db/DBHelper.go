package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DBhelper interface {
	GetMongoDBCollection(collectionName string) (*mongo.Collection, error)
}

type dbhelper struct {
	DatabaseName string
	client       *mongo.Client
	Logger       log.Logger
}

func NewDatabase(mongoURI string, mongoDatabase string, logger log.Logger) DBhelper {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected to MongoDB!")

	return dbhelper{
		client:       client,
		DatabaseName: mongoDatabase,
		Logger:       logger,
	}
}

func (d dbhelper) GetMongoDBCollection(collectionName string) (*mongo.Collection, error) {
	collection := d.client.
		Database(d.DatabaseName).
		Collection(collectionName)

	return collection, nil
}
