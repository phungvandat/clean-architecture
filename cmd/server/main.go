package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/joho/godotenv"
	"github.com/phungvandat/identity-service/endpoints"
	serviceHttp "github.com/phungvandat/identity-service/http"
	"github.com/phungvandat/identity-service/service"
	userSvc "github.com/phungvandat/identity-service/service/user"
	userRepo "github.com/phungvandat/identity-service/service/user/repository"
	userUseCase "github.com/phungvandat/identity-service/service/user/usecase"
	mongoDB "github.com/phungvandat/identity-service/util/config/db/mongo"
	envConfig "github.com/phungvandat/identity-service/util/config/env"
)

func main() {
	var err error
	if os.Getenv("ENV") == "local" {
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

	var h http.Handler
	{
		h = serviceHttp.NewHTTPHandler(
			endpoints.MakeServerEndpoints(s),
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

	logger.Log("exit", <-errs)
}
