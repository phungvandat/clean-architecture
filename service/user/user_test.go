package user

import (
	"context"
	"reflect"
	"testing"

	"github.com/phungvandat/clean-architecture/model/domain"
	userReq "github.com/phungvandat/clean-architecture/model/request/user"
	userRes "github.com/phungvandat/clean-architecture/model/response/user"
	repo "github.com/phungvandat/clean-architecture/repository"
	userRepo "github.com/phungvandat/clean-architecture/repository/user"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_userUsecase_FindByID(t *testing.T) {
	userRepoMock := new(userRepo.RepositoryMock)
	user := &domain.User{
		ID: primitive.NewObjectID(),
	}
	findByIDres := &userRes.FindByID{
		User: user,
	}
	userRepoMock.On("FindByID", mock.Anything, mock.Anything).Return(user, nil)
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
