package repository

import (
	"context"

	"github.com/phungvandat/identity-service/model/domain"
	"github.com/phungvandat/identity-service/util/constants"
	"github.com/phungvandat/identity-service/util/errors"
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
