package execution

import (
	"sync"
	"log"

	"github.com/gautamamber/mongo-to-es-golang/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Convert primitive mongo map
func convertPrimitiveMToMap(documents []primitive.M) []map[string]interface{} {
	var convertedDocuments []map[string]interface{}

	for _, doc := range documents {
		// Convert primitive.M to map[string]interface{}
		convertedDoc := make(map[string]interface{})
		for k, v := range doc {
			convertedDoc[k] = v
		}
		convertedDocuments = append(convertedDocuments, convertedDoc)
	}

	return convertedDocuments
}


func DumpDataMongoToES(collectionName string, ch chan<- string, wg *sync.WaitGroup) {

	// Ensure that the Done method is called at the end of the exection
	defer wg.Done()

	// Get Mongo data from specific mongo collection
	mongoDocuments, _ := utils.GetMongoDocuments(collectionName)
	convertedDocuments := convertPrimitiveMToMap(mongoDocuments)
	// Iterate in MongoDocument and Dump
	if err := utils.BulkDocumentAdd(convertedDocuments, collectionName); err != nil {
		log.Fatal("Error bulk indexing document")
	}
	ch <- collectionName + " - Dump successfully"
}
