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
	repo "github.com/phungvandat/clean-architecture/repository"
	userRepo "github.com/phungvandat/clean-architecture/repository/user"
	"github.com/phungvandat/clean-architecture/service"
	userSvc "github.com/phungvandat/clean-architecture/service/user"
	serviceGrpc "github.com/phungvandat/clean-architecture/transport/grpc"
	serviceHttp "github.com/phungvandat/clean-architecture/transport/http"
	mongoDB "github.com/phungvandat/clean-architecture/util/config/db/mongo"
	envConfig "github.com/phungvandat/clean-architecture/util/config/env"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	var err error
	var idProduction = os.Getenv("ENV") == "production"
	if !idProduction {
		err = godotenv.Load()
		if err != nil {
			panic(fmt.Sprintf("failed to load .env by error: %v", err))
		}
	}

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

	// Setup repository
	var (
		mongoDB, closeMongoDB = mongoDB.NewDB(envConfig.GetMogoDBName(), envConfig.GetMongoURL())
		userRepo              = userRepo.NewUserRepo(mongoDB)
		repo                  = repo.Repository{
			User: userRepo,
		}
	)

	// Setup service
	var (
		// user service
		userService = service.Compose(
			userSvc.NewUserService(repo),
			userSvc.ValidationMiddleware(),
		).(userSvc.Service)

		s = service.Service{
			UserService: userService,
		}
	)
	defer closeMongoDB()

	endpoints := endpoints.MakeServerEndpoints(s)

	// setup http
	httpPort := "3000"
	if envConfig.GetHttpPortEnv() != "" {
		httpPort = envConfig.GetHttpPortEnv()
	}
	httpAddr := fmt.Sprintf(":%v", httpPort)

	var httpHandler http.Handler
	{
		httpHandler = serviceHttp.NewHttpHandler(
			endpoints,
			logger,
		)
	}

	httpListener, err := net.Listen("tcp", httpAddr)
	if err != nil {
		panic(fmt.Sprintf("Create http listener failed by error: %v", err))
	}

	// setup grpc
	portGRPC := "4001"
	if envConfig.GetGrpcPortEnv() != "" {
		portGRPC = envConfig.GetGrpcPortEnv()
	}
	grpcAddr := fmt.Sprintf(":%v", portGRPC)

	opts := []grpc.ServerOption{}

	if os.Getenv("ENV") == "tls-secure" || idProduction {
		// Create the TLS credentials
		creds, err := tls.X509KeyPair([]byte(envConfig.GetServerCRT()), []byte(envConfig.GetServerKey()))
		if err != nil {
			logger.Log("could not load TLS keys", err)
		}
		httpTLSConfig := &tls.Config{
			Certificates: []tls.Certificate{creds},
		}

		httpListener = tls.NewListener(httpListener, httpTLSConfig)

		opts = append(opts, grpc.Creds(credentials.NewServerTLSFromCert(&creds)))
	}

	var (
		httpServer = http.Server{
			Handler: httpHandler,
		}

		grpcServer = grpc.NewServer(opts...)
	)

	serviceGrpc.NewGrpcHandler(
		endpoints,
		logger,
		grpcServer,
	)

	errChn := make(chan error)
	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		errChn <- fmt.Errorf("%s", <-ch)
	}()

	go func() {
		logger.Log("transport", "HTTP", "addr", httpAddr)
		errChn <- httpServer.Serve(httpListener)
	}()

	go func() {
		lis, err := net.Listen("tcp", grpcAddr)
		defer lis.Close()
		if err != nil {
			errChn <- err
		}
		logger.Log("transport", "GRPC", "addr", grpcAddr)
		errChn <- grpcServer.Serve(lis)
	}()

	logger.Log("exit", <-errChn)
}
