package middleware

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"runtime/debug"
)

func PanicRecovery() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		panicked := true

		defer func() {
			if r := recover(); r != nil || panicked {
				log.Println(string(debug.Stack()))
				err = status.Errorf(codes.Internal, "PANIC | request data : %v", req)
			}
		}()

		resp, err = handler(ctx, req)
		panicked = false
		return resp, err
	}
}
