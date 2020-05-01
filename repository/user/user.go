package user

import (
	"context"

	"github.com/phungvandat/clean-architecture/model/domain"
	"github.com/phungvandat/clean-architecture/util/constants"
	"github.com/phungvandat/clean-architecture/util/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	mongoDB *mongo.Database
}

// NewUserRepo create a new user repo
func NewUserRepo(mongoDB *mongo.Database) Repository {
	return &userRepo{
		mongoDB: mongoDB,
	}
}

// FindByID function handles find user with id condition from repository
func (repo *userRepo) FindByID(ctx context.Context, id string) (*domain.User, error) {
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
