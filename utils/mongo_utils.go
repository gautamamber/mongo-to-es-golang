package utils

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/gautamamber/mongo-to-es-golang/settings"
	"github.com/gautamamber/mongo-to-es-golang/connection"

)

func GetMongoDocuments(collectionName string) ([]bson.M, error) {
	
	mongoClientDB := MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := mongoClientDB.Collection(collectionName).Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var document bson.M
		if err == cursor.Decode(&document); err != nil {
			return nil, err
		}
		documents = append(documents, document)
	}

	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return documents, nil

}
