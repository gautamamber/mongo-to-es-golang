package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gautamamber/mongo-to-es-golang/connection"

)

func GetMongoDocuments(collectionName string) ([]bson.M, error) {
	
	mongoClientDB := connection.MongoDB
	// Use timer for context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// Cancel the context
	defer cancel()
	cursor, err := mongoClientDB.Collection(collectionName).Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var documents []bson.M

	for cursor.Next(ctx) {
		var document bson.M
		if err := cursor.Decode(&document); err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return documents, nil

}
