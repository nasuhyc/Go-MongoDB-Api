package config

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalln("Error loading .env.file")
	}

	mongoIRU := os.Getenv("MONGOURI")
	return mongoIRU

}

func ConnectDB() *mongo.Client {

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
	return client
}

var DB *mongo.Client = ConnectDB()

func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("employes").Collection(collectionName)
}
