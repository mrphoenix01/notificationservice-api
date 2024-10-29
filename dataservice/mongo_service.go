// dataservice/mongo_service.go
package dataservice

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// ConnectMongoDB initializes the MongoDB client
func ConnectMongoDB() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	Client = client
	log.Println("Connected to MongoDB")
}

// GetCollection returns a MongoDB collection from the specified database and collection name
func GetCollection(database, collection string) *mongo.Collection {
	return Client.Database(database).Collection(collection)
}
