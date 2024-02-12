package grpc

import (
	"fmt"
	"github.com/evenyosua18/auth2/app/constant"
	"github.com/evenyosua18/auth2/app/infrastructure/server/grpc/middleware"
	"github.com/evenyosua18/auth2/app/repository"
	"github.com/evenyosua18/auth2/app/utils/grpchelper"
	"github.com/evenyosua18/codes"
	"github.com/evenyosua18/sentry-helper"
	"github.com/evenyosua18/tracing"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunServer() {
	//setup environment variable (for local)
	if _, err := os.Stat(".env"); err == nil {
		if err := godotenv.Load(".env"); err != nil {
			panic(err)
		}
	}

	//setup grpc option
	var options []grpc.ServerOption
	options = append(options, grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     1 * time.Hour,
		MaxConnectionAge:      5 * time.Minute,
		MaxConnectionAgeGrace: 5 * time.Second,
	}))

	options = append(options, grpc.UnaryInterceptor(middleware.ChainUnaryServer(middleware.PanicRecovery(), middleware.GrpcLogger(), middleware.OauthClientValidation())))

	//initialize sentry
	flushFunction, err := sentry_helper.InitializeSentry(os.Getenv(constant.SentryDSN), os.Getenv(constant.AppEnv))
	if err != nil {
		panic(err)
	}
	defer flushFunction(os.Getenv(constant.SentryFlush))

	//setup tracer
	sentry_helper.SetRouter(&grpchelper.GrpcHelper{})
	sentry_helper.SetSkippedCaller(5, 3)
	sentry_helper.SetNamingRules(&grpchelper.ManageSentry{})
	tracing.SetTracer(sentry_helper.Get())
	tracing.ShowLog()

	//init codes
	codes.RegisterCode(os.Getenv(constant.CodePath))
	codes.SetUnknownCode(codes.Code{
		CustomCode:      999,
		ResponseMessage: "unknown error code",
		ErrorMessage:    "unknown error code",
		ResponseCode:    13,
	})

	//init connection
	repository.InitConnection()

	//create grpc server
	grpcServer := grpc.NewServer(options...)

	//register grpc server
	Apply(grpcServer)
	reflection.Register(grpcServer)

	//run grpc server
	go func() {
		lis, err := net.Listen("tcp", fmt.Sprintf(`%s:%s`, os.Getenv(constant.GrpcHost), os.Getenv(constant.GrpcPort)))

		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}

		if err = grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to start grpc server: %v", err)
		}
	}()

	log.Println(fmt.Sprintf("grpc server is running at %s:%s", os.Getenv(constant.GrpcHost), os.Getenv(constant.GrpcPort)))

	//get signal when server interrupted
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-c
	log.Fatalf("process killed with signal: %s", sig.String())
}
