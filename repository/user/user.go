package user

import (
	"context"

	"github.com/phungvandat/clean-architecture/model/domain"
	"github.com/phungvandat/clean-architecture/model/repository"
	userInput "github.com/phungvandat/clean-architecture/model/repository/user"
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
func (repo *userRepo) FindByID(ctx context.Context, id string, options ...*repository.RepoOptions) (*domain.User, error) {
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

// Find function handles find users from repository with conditions
func (repo *userRepo) Find(ctx context.Context, conditions userInput.FindConditions, options ...*repository.RepoOptions) ([]*domain.User, error) {
	var (
		userCollection = repo.mongoDB.Collection(constants.MongoUserCollection)
		users          = []*domain.User{}
	)

	matchMap := bson.M{}
	if conditions.Fullname != "" {
		matchMap["fullname"] = conditions.Fullname
	}

	matchPl := bson.M{"$match": matchMap}

	pipeline := []bson.M{
		matchPl,
	}

	cursor, err := userCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		user := &domain.User{}
		err = cursor.Decode(user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Create function handles create user
func (repo *userRepo) Create(ctx context.Context, user *domain.User, options ...*repository.RepoOptions) (*domain.User, error) {
	repoOptions := repository.MergeRepoOptions(options...)
	if repoOptions.TX != nil {
		ctx = repoOptions.TX.SCtx
	}

	userCollection := repo.mongoDB.Collection(constants.MongoUserCollection)
	insertResult, err := userCollection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = insertResult.InsertedID.(primitive.ObjectID)

	return user, nil
}
