package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"crypto/tls"

	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
	"github.com/phungvandat/clean-architecture/endpoints"
	"github.com/phungvandat/clean-architecture/service"
	userSvc "github.com/phungvandat/clean-architecture/service/user"
	userRepo "github.com/phungvandat/clean-architecture/service/user/repository"
	userUseCase "github.com/phungvandat/clean-architecture/service/user/usecase"
	serviceGrpc "github.com/phungvandat/clean-architecture/transport/grpc"
	serviceHttp "github.com/phungvandat/clean-architecture/transport/http"
	mongoDB "github.com/phungvandat/clean-architecture/util/config/db/mongo"
	envConfig "github.com/phungvandat/clean-architecture/util/config/env"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var err error
	if os.Getenv("ENV") != "production" {
		err = godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("failed to load .env by error: %v", err))
		}
	}

	// Setup addr
	port := "3000"
	if envConfig.GetPortEnv() != "" {
		port = envConfig.GetPortEnv()
	}

	httpAddr := fmt.Sprintf(":%v", port)

	// Setup log
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	// Setup locale
	{
		loc, err := time.LoadLocation("Asia/Bangkok")
		if err != nil {
			logger.Log("error", err)
			os.Exit(1)
		}
		time.Local = loc
	}

	// Setup service
	var (
		mongoDB, closeMongoDB = mongoDB.NewDB(envConfig.GetMogoDBName(), envConfig.GetMongoURI())

		// user service
		userRepo    = userRepo.NewUserRepo(mongoDB)
		userService = service.Compose(
			userUseCase.NewUserUseCase(userRepo),
			userUseCase.ValidationMiddleware(),
		).(userSvc.Service)

		s = service.Service{
			UserService: userService,
		}
	)
	defer closeMongoDB()

	endpoints := endpoints.MakeServerEndpoints(s)

	var h http.Handler
	{
		h = serviceHttp.NewHTTPHandler(
			endpoints,
			logger,
		)
	}

	errs := make(chan error)
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-ch)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", httpAddr)
		errs <- http.ListenAndServe(httpAddr, h)

	}()

	// grpc server
	portGRPC := "4001"
	if envConfig.GetGRPCPortEnv() != "" {
		portGRPC = envConfig.GetGRPCPortEnv()
	}

	// Create an array of gRPC options with the credentials
	opts := []grpc.ServerOption{}

	if os.Getenv("ENV") == "secure-grpc" {
		// Create the TLS credentials
		creds, err := tls.X509KeyPair([]byte(envConfig.GetServerPem()), []byte(envConfig.GetServerKey()))
		if err != nil {
			logger.Log("could not load TLS keys", err)
		}
		opts = append(opts, grpc.Creds(credentials.NewServerTLSFromCert(&creds)))
	}

	var (
		grpcServer = grpc.NewServer(opts...)
		grpcAddr   = fmt.Sprintf(":%v", portGRPC)
	)
	serviceGrpc.NewGRPCHandler(
		endpoints,
		logger,
		grpcServer,
	)

	go func() {
		lis, err := net.Listen("tcp", grpcAddr)
		defer lis.Close()
		if err != nil {
			errs <- err
		}
		logger.Log("transport", "GRPC", "addr", grpcAddr)
		errs <- grpcServer.Serve(lis)
	}()

	logger.Log("exit", <-errs)
}
