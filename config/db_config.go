package config

import (
	"os"
	"fmt"
)
type MongoConfig struct {
	MONGO_HOST    string
	MONGO_PORT    string
	MONGO_DB_NAME string
}


type ElasticsearchConfig struct {
	ES_HOST        string
	ES_PORT        string
	INDEX_NAME     string
	INDEX_PREFIX   string
}

// Get Mongo config method
func GetMongoConfig() *MongoConfig{

	fmt.Println("Fetching Mongo Configs....")

	config := &MongoConfig{
		MONGO_HOST:    os.Getenv("MONGO_HOST"),
		MONGO_PORT:    os.Getenv("MONGO_PORT"),
		MONGO_DB_NAME: os.Getenv("MONGO_DB_NAME"),
	}

	return config
}


// Get ES config method
func getEsConfig() *ElasticsearchConfig{

	fmt.Println("Fetching ES Configs....")
	config := &ElasticsearchConfig{
		ES_HOST:      os.Getenv("ES_HOST"),
		ES_PORT:      os.Getenv("ES_PORT"),
		INDEX_NAME:   os.Getenv("ES_INDEX_NAME"),
		INDEX_PREFIX: os.Getenv("ES_INDEX_PREFIX"),
	}

	return config
}
