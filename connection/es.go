package connection

import (
	"context"
	"log"

	"github.com/gautamamber/mongo-to-es-golang/config"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
)

var (
	// Global ES connection
	ElasticsearchClient *elasticsearch.Client
)

func InitElasticSearch(ctx context.Context) error {
	var err error

	esConfig := config.GetEsConfig()

	ElasticsearchClient, err = elasticsearch.NewClient(elasticsearch.Config{
        Addresses: []string{
            esConfig.ES_HOST,
        },
        Username: esConfig.ES_USERNAME,
        Password: esConfig.ES_PASSWORD,
    })

	if err != nil {
		log.Fatalf("Error creating Elasticsearch client: %s", err)
		panic(err)
	}

	// Check if client is working by pinging ES
	res, err := ElasticsearchClient.Info()
	if err != nil {
		log.Fatalf("Error getting Elasticsearch info: %s", err)
	}
	defer res.Body.Close()
	// Print the status code and response
    log.Printf("Elasticsearch client created successfully: %s", res.Status())

	return nil
}