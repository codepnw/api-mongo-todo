package databases

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func ConnectMongo() (*mongo.Client, error) {
	dburi := os.Getenv("DB_URI")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")

	clientOptions := options.Client().ApplyURI(dburi)
	clientOptions.SetAuth(options.Credential{
		Username: username,
		Password: password,
	})
	
	connection, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	client = connection

	log.Println("connected to mongodb...")
	return client, nil
}

func GetClient() *mongo.Client {
	return client
}

func GetTodosCollection() *mongo.Collection {
	return client.Database(os.Getenv("DB_NAME")).Collection("todos")
}
