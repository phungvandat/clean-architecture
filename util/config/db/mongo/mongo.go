package mongo

import (
	"context"
	"fmt"

	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewDB(dbName string, uri string) (*mongo.Database, func()) {
	clientOptions := options.Client().ApplyURI(uri)
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic(err)
	}

	db := client.Database(dbName)

	return db, func() {
		err := client.Disconnect(ctx)

		if err != nil {
			log.Println("Failed to close DB by error: ", err)
			return
		}
		fmt.Println("Closed DB")
	}
}
