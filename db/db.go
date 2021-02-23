package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoDB is
func MongoDB() (*mongo.Database, error) {
	clientOptions := options.Client()
	// credential := options.Credential{
	// 	Username: "admin",
	// 	Password: "admin",
	// }
	clientOptions.ApplyURI("mongodb://localhost:27017")
	// clientOptions.ApplyURI("mongodb://mongodb").SetAuth(credential)
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.Background())
	if err != nil {
		return nil, err
	}

	return client.Database("simpus"), nil
}
