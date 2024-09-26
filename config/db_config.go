package config

import (
	"os"
)
type MongoConfig struct {
	MONGO_HOST    string
	MONGO_PORT    string
	MONGO_DB_NAME string
}


type ElasticsearchConfig struct {
	ES_HOST        string
	ES_USERNAME    string
	ES_PASSWORD    string
	INDEX_NAME     string
	INDEX_PREFIX   string
}

// Get Mongo config method
func GetMongoConfig() *MongoConfig{

	config := &MongoConfig{
		MONGO_HOST:    os.Getenv("MONGO_HOST"),
		MONGO_PORT:    os.Getenv("MONGO_PORT"),
		MONGO_DB_NAME: os.Getenv("MONGO_DB_NAME"),
	}

	return config
}


// Get ES config method
func GetEsConfig() *ElasticsearchConfig{

	config := &ElasticsearchConfig{
		ES_HOST:      os.Getenv("ES_HOST"),
		ES_USERNAME:  os.Getenv("ES_USERNAME"),
		ES_PASSWORD:  os.Getenv("ES_PASSWORD"),
		INDEX_NAME:   os.Getenv("ES_INDEX_NAME"),
		INDEX_PREFIX: os.Getenv("ES_INDEX_PREFIX"),
	}

	return config
}
