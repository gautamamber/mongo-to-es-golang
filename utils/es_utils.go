package utils

import (
	"fmt"
	"log"
	"context"
	"strings"

	"github.com/gautamamber/mongo-to-es-golang/config"
	"github.com/gautamamber/mongo-to-es-golang/connection"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)


func CreateIndexAndMapping() {

	esConfig := config.GetEsConfig()
	
	indexName := esConfig.INDEX_PREFIX + "_" + esConfig.INDEX_NAME

	// Create Index, but before this check if index already exists, if so delete and then create again

	res, err := connection.ElasticsearchClient.Indices.Exists([]string{indexName})

	if err != nil {
		log.Fatalf("Error while checking if index already exists: %s", err)
	}

	if res.StatusCode == 200 {
		fmt.Printf("Index '%s' already exists\n", indexName)
		// Delete Index
		req := esapi.IndicesDeleteRequest{
			Index: []string{indexName},
		}
		res, err := req.Do(context.Background(), connection.ElasticsearchClient)
		if err != nil {
			log.Fatalf("Error deleting index '%s': %s", indexName, err)
		}
		defer res.Body.Close()
		if res.IsError() {
			log.Printf("Error deleting index '%s': %s", indexName, res.String())
		} else {
			fmt.Printf("Index '%s' deleted successfully.\n", indexName)
		}
	}
	
	// Define mapping

	var Mapping = `{
		"mappings": {
			"properties": {
				"doc_type": {
					"type": "keyword"
				},
				"name": {
					"type": "text"
				},
				"value": {
					"type": "float"
				}
			}
		}
	}`

	// Create the index with the provided mapping
	req := esapi.IndicesCreateRequest{
		Index: indexName,
		Body:  strings.NewReader(Mapping),
	}

	response, err := req.Do(context.Background(), connection.ElasticsearchClient)
	if err != nil {
		log.Fatalf("Error creating index '%s': %s", indexName, err)
	}
	defer response.Body.Close()

	fmt.Printf("Index '%s' created.\n", indexName)

}