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
	"sync"

	"github.com/gautamamber/mongo-to-es-golang/settings"
	"github.com/gautamamber/mongo-to-es-golang/connection"
	"github.com/gautamamber/mongo-to-es-golang/execution"

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

	// Initialize ES client
	err = connection.InitElasticSearch(ctx)
	if err != nil {
		log.Fatal("Error Connecting ES:", err.Error())
	}
}

func main() {
	// Script execution begins here

	fmt.Println("Script Execution begins...")

	// Get List of All existing collections
	collections := settings.GetListOfStrings()

	// Create channel for communication

	ch := make(chan string)

	// Create a wait group to wait for all goroutine to finish
	var wg sync.WaitGroup
	
	for _, collection := range collections {
		// Increment the wait group counter before launching go routine
		wg.Add(1)
		go execution.DumpDataMongoToES(collection, ch, &wg)
	}

	// Launch a go routine to close channel once all the collections are processed
	go func() {
		wg.Wait()
		close(ch)
	}()
	// Print message from channel
	for msg := range ch {
		fmt.Println("Message in Channel respone" + msg)
	}

	fmt.Println("Script Executed Successfully!")

}
