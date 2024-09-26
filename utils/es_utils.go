package utils

import (
	"fmt"
	"log"
	"encoding/json"
	"context"
	"strings"
	"bytes"

	"github.com/gautamamber/mongo-to-es-golang/config"
	"github.com/gautamamber/mongo-to-es-golang/connection"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// Document structure with only relevant fields
type Document struct {
	DocType string  `json:"doc_type"`
	Name    string  `json:"name"`
	Value   string `json:"value"`
}

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
					"type": "string"
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

func BulkDocumentAdd(documents []map[string]interface{}, collectionName string) error {

	esConfig := config.GetEsConfig()
	
	indexName := esConfig.INDEX_PREFIX + "_" + esConfig.INDEX_NAME

	var buf bytes.Buffer 

	// Prepare the bulk request

	for _, doc := range documents {

		// Extract only the needed fields
		name, okName := doc["name"].(string)
		value, okValue := doc["value"].(string)

		// Skip the document if any of the required fields are missing
		if !okName || !okValue {
			continue
		}

		// Create a new Document with only the necessary fields
		esDoc := Document{
			DocType: collectionName,
			Name:    name,
			Value:   value,
		}
		meta := map[string]interface{}{
			"index": map[string]interface{}{
				"_index": indexName,
			},
		}
		// Write metadata to buffer
		if err := json.NewEncoder(&buf).Encode(meta); err != nil {
			return fmt.Errorf("failed to encode metadata: %w", err)
		}
		// Write the actual document to buffer
		if err := json.NewEncoder(&buf).Encode(esDoc); err != nil {
			return fmt.Errorf("failed to encode document: %w", err)
		}
	}
	req := esapi.BulkRequest{
		Body: &buf,
		Refresh: "true", // Make the documents available for search immediately
	}

	// Execute the bulk request
	res, err := req.Do(context.Background(), connection.ElasticsearchClient)
	if err != nil {
		return fmt.Errorf("failed to execute bulk request: %w", err)
	}
	defer res.Body.Close()
	if res.IsError() {
		return fmt.Errorf("error indexing documents: %s", res.String())
	}

	fmt.Println("Documents indexed successfully.")
	return nil
}
