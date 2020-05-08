package mongo

import (
	"context"
	"fmt"

	"log"

	"github.com/phungvandat/clean-architecture/util/constants"
	"github.com/phungvandat/clean-architecture/util/helper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// NewDB func return a mongo connection
func NewDB(dbName string, url string) (*mongo.Database, func()) {
	clientOptions := options.Client().ApplyURI(url)
	ctx := context.TODO()
	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}

	log.Println("Connected db")

	db := client.Database(dbName)

	return db, func() {
		err := client.Disconnect(ctx)

		if err != nil {
			log.Println("Failed to close DB by error: ", err)
			return
		}
		fmt.Println("Closed connection DB")
	}
}

// PrepareDB function to create collection if not exists in db
func PrepareDB(mDB *mongo.Database) {
	var (
		err error
		ctx = context.TODO()
	)
	collNames, err := mDB.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	requireCollNames := []string{
		constants.MongoUserCollection,
	}
	arr, err := helper.ConvertTypeArrayToInterfaceArray(collNames)
	if err != nil {
		panic(err)
	}
	for _, n := range requireCollNames {
		if !helper.CheckArrIncludeItem(arr, n) {
			err = mDB.CreateCollection(ctx, n)
			if err != nil {
				panic(err)
			}
		}
	}
}
