package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/codepnw/go-mongo-todos/databases"
	"github.com/codepnw/go-mongo-todos/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var envFile = "dev.env"

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := godotenv.Load(envFile); err != nil {
		log.Panicln("loading env failed", err)
	}

	// Database Mongo
	mongoClient, err := databases.ConnectMongo()
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := gin.Default()
	routers.NewRouter(app)

	app.Run(":" + os.Getenv("APP_PORT"))
}
