package user

import (
	"context"

	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"github.com/phungvandat/clean-architecture/endpoints"
	userDecode "github.com/phungvandat/clean-architecture/transport/grpc/decode/user"
	userEncode "github.com/phungvandat/clean-architecture/transport/grpc/encode/user"
	userproto "github.com/phungvandat/clean-architecture/transport/grpc/proto/user"
)

type userServer struct {
	findByID grpcTransport.Handler
}

func UserServer(
	endpoints endpoints.Endpoints,
	options []grpcTransport.ServerOption,
) userproto.UserServiceServer {
	return &userServer{
		findByID: grpcTransport.NewServer(
			endpoints.UserEndpoint.FindByID,
			userDecode.FindByID,
			userEncode.FindByID,
			options...,
		),
	}
}

func (s *userServer) FindByID(ctx context.Context, req *userproto.FindByIDRequest) (*userproto.FindByIDResponse, error) {
	_, res, err := s.findByID.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*userproto.FindByIDResponse), nil
}
