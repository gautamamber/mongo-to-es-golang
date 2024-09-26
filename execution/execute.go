package execution

import (
	"sync"
	"fmt"

	"github.com/gautamamber/mongo-to-es-golang/utils"
)

func DumpDataMongoToES(collectionName string, ch chan<- string, wg *sync.WaitGroup) {

	// Ensure that the Done method is called at the end of the exection
	defer wg.Done()

	// Get Mongo data from specific mongo collection
	mongoDocument, _ := utils.GetMongoDocuments(collectionName)
	fmt.Println(mongoDocument)
	// Dump in ES

	// Return to channel

	data := "Data from " + collectionName

	fmt.Println(data)
	ch <- collectionName + "Dump successfully"

}