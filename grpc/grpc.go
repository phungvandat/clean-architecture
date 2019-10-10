package gprc

import (
	"github.com/go-kit/kit/log"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"github.com/phungvandat/clean-architecture/endpoints"
	"github.com/phungvandat/clean-architecture/grpc/options"
	userproto "github.com/phungvandat/clean-architecture/grpc/proto/user"
	userServer "github.com/phungvandat/clean-architecture/grpc/server/user"
	"google.golang.org/grpc"
)

func NewGRPCHandler(
	endpoints endpoints.Endpoints,
	logger log.Logger,
	grpcServer *grpc.Server,
) {
	options := []grpcTransport.ServerOption{
		grpcTransport.ServerBefore(options.LogRequestInfo(logger)),
		grpcTransport.ServerErrorLogger(logger),
	}

	userSvr := userServer.UserServer(endpoints, options)

	userproto.RegisterUserServiceServer(grpcServer, userSvr)
}
