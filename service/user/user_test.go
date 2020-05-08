package user

import (
	"context"
	"reflect"
	"testing"

	"github.com/phungvandat/clean-architecture/model/domain"
	repoOptions "github.com/phungvandat/clean-architecture/model/repository"
	userInput "github.com/phungvandat/clean-architecture/model/repository/user"
	userReq "github.com/phungvandat/clean-architecture/model/request/user"
	userRes "github.com/phungvandat/clean-architecture/model/response/user"
	repo "github.com/phungvandat/clean-architecture/repository"
	userRepo "github.com/phungvandat/clean-architecture/repository/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_userUsecase_FindByID(t *testing.T) {
	t.Parallel()
	user := &domain.User{
		ID: primitive.NewObjectID(),
	}
	findByIDres := &userRes.FindByID{
		User: user,
	}
	userRepoMock := &userRepo.RepositoryMock{
		FindByIDFunc: func(ctx context.Context, id string, options ...*repoOptions.RepoOptions) (*domain.User, error) {
			return user, nil
		},
	}
	repoMock := repo.NewRepositoryMock(repo.RepositoryMock{
		User: userRepoMock,
	})

	type args struct {
		ctx context.Context
		req userReq.FindByID
	}
	tests := []struct {
		name    string
		args    args
		want    *userRes.FindByID
		wantErr error
	}{
		{
			name: "Find user by ID success",
			args: args{
				ctx: context.TODO(),
				req: userReq.FindByID{
					UserID: user.ID.String(),
				},
			},
			want:    findByIDres,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			userSvc := &userService{
				repo: repoMock,
			}
			got, err := userSvc.FindByID(tt.args.ctx, tt.args.req)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("userUsecase.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userUsecase.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_Find(t *testing.T) {
	t.Parallel()
	userRepoMock := &userRepo.RepositoryMock{
		FindFunc: func(ctx context.Context, conditions userInput.FindConditions, options ...*repoOptions.RepoOptions) ([]*domain.User, error) {
			return []*domain.User{}, nil
		},
	}
	repoMock := repo.NewRepositoryMock(repo.RepositoryMock{
		User: userRepoMock,
	})
	type args struct {
		ctx context.Context
		req userReq.Find
	}
	tests := []struct {
		name    string
		args    args
		want    *userRes.Find
		wantErr error
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{
				repo: repoMock,
			}
			got, err := s.Find(tt.args.ctx, tt.args.req)
			if err != nil && err != tt.wantErr {
				t.Errorf("userService.Find() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_userService_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		req userReq.Create
	}
	tests := []struct {
		name    string
		args    args
		want    *userRes.Create
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &userService{}
			got, err := s.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("userService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
