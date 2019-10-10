package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/phungvandat/clean-architecture/model/domain"
	"github.com/phungvandat/clean-architecture/util/constants"
	"github.com/phungvandat/clean-architecture/util/engine"
	"github.com/phungvandat/clean-architecture/util/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoUserRepo struct {
	mongoDB *mongo.Database
}

// NewMongoUserRepo create a new user repo
func NewUserRepo(mongoDB *mongo.Database) Repository {
	return &mongoUserRepo{
		mongoDB: mongoDB,
	}
}

func (repo *mongoUserRepo) FindByID(ctx context.Context, id string) (*domain.User, error) {
	user := &domain.User{}
	userCollection := repo.mongoDB.Collection(constants.MongoUserCollection)
	objectID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{
		{"_id", objectID},
	}

	result := userCollection.FindOne(ctx, filter)

	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return nil, errors.UserNotExistError
		}
		return nil, result.Err()
	}

	err := result.Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *mongoUserRepo) TestAddTranslateQuery(ctx context.Context, query *engine.Query) ([]*domain.User, error) {
	users := []*domain.User{}
	userCollection := repo.mongoDB.Collection(constants.MongoUserCollection)

	filter := engine.TranslateQueryToMongoFilter(query)
	x, _ := json.Marshal(filter)
	fmt.Println(string(x))
	cursors, err := userCollection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cursors.Next(ctx) {
		user := &domain.User{}
		err = cursors.Decode(user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
