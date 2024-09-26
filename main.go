package main

// Objective: Migrate MongoDB collections in Elasticsearch Index
// Flow:
// Initialize and get the env variables the env variables
// Create Mongo Connection - During script initialization
	// Check if auth enabled then create authenticated connection else without auth
// Create ES Connection - During script initialization
	// Check if auth enabled then create authenticated connection else without auth
	
// Create Struct for ES config

// ES Flow:
// Check if same name index already exists of not, if exists then first delete
// Create new index with prefix_indexname (provided in .env)
// Create new mapping
// Create utility to Push data in ES

// Mongo utils
// Define list of collections
// Create method to fetch the documents from collections on the basis of collection name

// In main script runner utility
// Create go anonymous function for list of collections
// Get the data from Mongo for all the collection, map the data according to ES mapping and push in ES


import (
	"context"
	"fmt"
	"log"
	// "time"
	"github.com/gautamamber/mongo-to-es-golang/settings"
	"github.com/gautamamber/mongo-to-es-golang/connection"
	// "github.com/gautamamber/mongo-to-es-golang/utils"


	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
	// "github.com/elastic/go-elasticsearch/v7"
	// "github.com/elastic/go-elasticsearch/v7/esapi"
)

func init() {

	var err error
	ctx := context.Background()
	settings.InitEnvironmnetVars()
	
	// Initialize Mongo connection
	err = connection.InitMongo(ctx)
	if err != nil {
		log.Fatal("Error Connecting Mongo:", err.Error())
	}

}

func main() {
	// Script execution begins here

	fmt.Println("Script Execution begins...")

}