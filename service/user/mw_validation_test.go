package user

import (
	"context"
	"reflect"
	"testing"

	userReq "github.com/phungvandat/clean-architecture/model/request/user"
	userRes "github.com/phungvandat/clean-architecture/model/response/user"
	"github.com/phungvandat/clean-architecture/util/errors"
)

func Test_validationMiddleware_FindByID(t *testing.T) {
	t.Parallel()
	useRes := &userRes.FindByID{}
	userSvcMock := &ServiceMock{
		FindByIDFunc: func(ctx context.Context, req userReq.FindByID) (*userRes.FindByID, error) {
			return useRes, nil
		},
	}

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
			name: "Valid find by ID",
			args: args{
				ctx: context.TODO(),
				req: userReq.FindByID{
					UserID: "5c127a068eda730c3516b07f",
				},
			},
			want:    useRes,
			wantErr: nil,
		},
		{
			name: "User ID is required error",
			args: args{
				ctx: context.TODO(),
				req: userReq.FindByID{},
			},
			want:    nil,
			wantErr: errors.UserIDIsRequiredError,
		},
		{
			name: "Incorrect type of user ID error",
			args: args{
				ctx: context.TODO(),
				req: userReq.FindByID{
					UserID: "xxxx",
				},
			},
			want:    nil,
			wantErr: errors.IncorrectTypeOfUserIDError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mw := validationMiddleware{
				Service: userSvcMock,
			}
			got, err := mw.FindByID(tt.args.ctx, tt.args.req)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("validationMiddleware.FindByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("validationMiddleware.FindByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
