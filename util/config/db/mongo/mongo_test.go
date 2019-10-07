package mongo

// const (
// 	mongURI = "mongodb://localhost:27017"
// 	dbName  = "test"
// )

// func CreateTestDatabase(t *testing.T) (*mongo.Database, func()) {
// 	testURI := fmt.Sprintf("%s", mongURI)
// 	testDBName := fmt.Sprintf("%s", dbName)

// 	if os.Getenv("MONGO_TESTING_URI") != "" {
// 		testURI = os.Getenv("MONGO_TESTING_URI")
// 	}

// 	if os.Getenv("MONGO_TESING_DB_NAME") != "" {
// 		testDBName = os.Getenv("MONGO_TESING_DB_NAME")
// 	}

// 	clientOptions := options.Client().ApplyURI(testURI)
// 	ctx := context.TODO()
// 	client, err := mongo.Connect(ctx, clientOptions)

// 	if err != nil {
// 		t.Fatalf("Fail to connect database to testing by error %s", err.Error())
// 	}

// 	dbTest := client.Database(testDBName)
// 	return dbTest, func() {
// 		collections, err := dbTest.ListCollectionNames(ctx, nil)
// 		if err != nil {
// 			t.Fatalf("Fail to get all collection by error %s", err.Error())
// 		}
// 		for _, collectionName := range collections {
// 			err = dbTest.Collection(collectionName).Drop(ctx)
// 			if err != nil {
// 				t.Fatalf("Fail to drop collection %s by error %s", collectionName, err.Error())
// 			}
// 		}
// 		err = client.Disconnect(ctx)
// 		if err != nil {
// 			t.Fatalf("Fail to disconnect testing database by error %s", err.Error())
// 		}
// 	}
// }
