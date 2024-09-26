package connection

import (
	"context"
	"fmt"
	"log"

	"github.com/gautamamber/mongo-to-es-golang/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	// Global database connection pool
	MongoDB *mongo.Database
)

func MongoURI() string {
	// Fetch Mongo Config
	mongoConfig := config.GetMongoConfig()

	return fmt.Sprintf("mongodb://%s:%s/%s",
		mongoConfig.MONGO_HOST, mongoConfig.MONGO_PORT, mongoConfig.MONGO_DB_NAME)
}


func InitMongo(ctx context.Context) error {
	var err error

	mongoURI := MongoURI()

	fmt.Println("Mongo URI: ", mongoURI)
	// Validate if Mongo URI is empty
	// TODO: Validate Mongo URI using regex
	if mongoURI == "" {
		log.Fatal("Invalid Mongo config")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Create Mongo client
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
		panic(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
		panic(err)
	}
	mongoDBName := config.GetMongoConfig().MONGO_DB_NAME
	MongoDB = client.Database(mongoDBName)

	log.Println("Connected to MongoDB successfully!")
	return nil

}
