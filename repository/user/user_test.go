package user

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/phungvandat/clean-architecture/model/domain"
	"github.com/phungvandat/clean-architecture/model/repository"
	userInput "github.com/phungvandat/clean-architecture/model/repository/user"
	mongoConfig "github.com/phungvandat/clean-architecture/util/config/db/mongo"
	"github.com/phungvandat/clean-architecture/util/constants"
	"github.com/phungvandat/clean-architecture/util/errors"
	"github.com/phungvandat/clean-architecture/util/helper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_userRepo_FindByID(t *testing.T) {
	t.Parallel()
	dbTest, cleanup := mongoConfig.CreateTestDatabase(t)
	defer cleanup()

	userCollection := dbTest.Collection(constants.MongoUserCollection)
	user := &domain.User{
		ID:       primitive.NewObjectID(),
		Username: "cleanarch",
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

func Test_userRepo_Find(t *testing.T) {
	t.Parallel()
	dbTest, cleanup := mongoConfig.CreateTestDatabase(t)
	defer cleanup()

	var (
		ctx            = context.TODO()
		userCollection = dbTest.Collection(constants.MongoUserCollection)
		user1          = domain.User{
			ID:       primitive.NewObjectID(),
			Fullname: "cleanarch",
		}
		user2 = domain.User{
			ID:       primitive.NewObjectID(),
			Fullname: "cleanarch",
		}
		user3 = domain.User{
			ID:       primitive.NewObjectID(),
			Fullname: "pvd",
		}
	)

	insertArr, _ := helper.ConvertTypeArrayToInterfaceArray([]domain.User{user1, user2, user3})

	_, err := userCollection.InsertMany(ctx, insertArr)
	if err != nil {
		t.Fatalf("Insert user document failed by error %v", err)
	}

	type args struct {
		ctx        context.Context
		conditions userInput.FindConditions
	}

	tests := []struct {
		name    string
		args    args
		want    []*domain.User
		wantErr error
	}{
		{
			name: "Find with conditions success",
			args: args{
				ctx:        ctx,
				conditions: userInput.FindConditions{},
			},
			want:    []*domain.User{&user1, &user2, &user3},
			wantErr: nil,
		},
		{
			name: "Find with fullname conditions success",
			args: args{
				ctx: ctx,
				conditions: userInput.FindConditions{
					Fullname: "cleanarch",
				},
			},
			want:    []*domain.User{&user1, &user2},
			wantErr: nil,
		},
		{
			name: "Find with fullname conditions not match",
			args: args{
				ctx: ctx,
				conditions: userInput.FindConditions{
					Fullname: "cleanarch2",
				},
			},
			want:    []*domain.User{},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &userRepo{
				mongoDB: dbTest,
			}

			got, err := repo.Find(tt.args.ctx, tt.args.conditions)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("userRepo.Find() error = %v, wantErr = %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.Find() = %v, want = %v", got, tt.want)
			}
		})
	}

}

func Test_userRepo_Create(t *testing.T) {
	t.Parallel()
	dbTest, cleanup := mongoConfig.CreateTestDatabase(t)
	defer cleanup()

	user := &domain.User{
		ID:        primitive.NewObjectID(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  "ca",
	}

	type args struct {
		ctx     context.Context
		user    *domain.User
		options []*repository.RepoOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.User
		wantErr error
	}{
		{
			name: "Create user success",
			args: args{
				ctx:  context.TODO(),
				user: user,
			},
			want: user,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &userRepo{
				mongoDB: dbTest,
			}
			got, err := repo.Create(tt.args.ctx, tt.args.user, tt.args.options...)
			if err != nil && err != tt.wantErr {
				t.Errorf("userRepo.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepo.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
