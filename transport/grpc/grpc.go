package gprc

import (
	"github.com/go-kit/kit/log"
	grpcTransport "github.com/go-kit/kit/transport/grpc"
	"github.com/phungvandat/clean-architecture/endpoints"
	"github.com/phungvandat/clean-architecture/transport/grpc/options"
	userproto "github.com/phungvandat/clean-architecture/transport/grpc/proto/user"
	userServer "github.com/phungvandat/clean-architecture/transport/grpc/server/user"
	"google.golang.org/grpc"
)

func NewGrpcHandler(
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
