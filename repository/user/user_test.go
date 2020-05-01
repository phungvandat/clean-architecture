package user

import (
	"context"
	"reflect"
	"testing"

	"github.com/phungvandat/clean-architecture/model/domain"
	mongoConfig "github.com/phungvandat/clean-architecture/util/config/db/mongo"
	"github.com/phungvandat/clean-architecture/util/constants"
	"github.com/phungvandat/clean-architecture/util/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_userRepo_FindByID(t *testing.T) {
	t.Parallel()
	dbTest, cleanup := mongoConfig.CreateTestDatabase(t)
	defer cleanup()

	userCollection := dbTest.Collection(constants.MongoUserCollection)
	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Username: "phungvandat",
	}

	_, err := userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		t.Fatalf("Insert user document failed by error %v", err)
	}

	type args struct {
		ctx context.Context
		id  string
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.User
		wantErr error
	}{
		{
			name: "Find user by ID success",
			args: args{
				ctx: context.TODO(),
				id:  user.ID.Hex(),
			},
			want:    user,
			wantErr: nil,
		},
		{
			name: "Find user by ID failed by not exist",
			args: args{
				ctx: context.TODO(),
				id:  "5c127a068eda730c3516b07f",
			},
			want:    nil,
			wantErr: errors.UserNotExistError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &userRepo{
				mongoDB: dbTest,
			}
			got, err := repo.FindByID(tt.args.ctx, tt.args.id)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("userRepo.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
