package middleware

import (
	"context"
	sentry_helper "github.com/evenyosua18/ego-util/tracing/sentry-helper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func GrpcLogger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (response interface{}, err error) {
		log.Printf("%s\n", info.FullMethod)

		response, err = handler(ctx, req)

		if err != nil && status.Code(err) == codes.Internal {
			log.Printf("internal error: %s - %v\n", info.FullMethod, err)
			sentry_helper.AlertError(err, map[string]string{
				"service": info.FullMethod,
			})
		}

		return
	}
}
