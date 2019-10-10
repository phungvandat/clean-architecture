package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/phungvandat/clean-architecture/model/domain"
	"github.com/phungvandat/clean-architecture/model/request"
	"github.com/phungvandat/clean-architecture/model/response"
	userRepo "github.com/phungvandat/clean-architecture/service/user/repository"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Test_userUsecase_FindByID(t *testing.T) {
	userRepoMock := new(userRepo.RepositoryMock)
	userRes := &domain.User{
		ID: primitive.NewObjectID(),
	}
	res := &response.FindByID{
		User: userRes,
	}
	userRepoMock.On("FindByID", mock.Anything, mock.Anything).Return(userRes, nil)

	type args struct {
		ctx context.Context
		req request.FindByID
	}
	tests := []struct {
		name    string
		args    args
		want    *response.FindByID
		wantErr error
	}{
		{
			name: "Find user by ID success",
			args: args{
				ctx: context.TODO(),
				req: request.FindByID{
					UserID: userRes.ID.Hex(),
				},
			},
			want:    res,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			useCase := &userUsecase{
				userRepo: userRepoMock,
			}
			got, err := useCase.FindByID(tt.args.ctx, tt.args.req)
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
