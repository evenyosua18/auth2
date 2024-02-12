package middleware

import (
	"context"
	"github.com/evenyosua18/auth2/app/infrastructure/container"
	"github.com/evenyosua18/auth2/app/repository"
	"github.com/evenyosua18/auth2/app/utils/grpchelper"
	"github.com/evenyosua18/auth2/app/utils/response"
	"github.com/evenyosua18/tracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

const (
	id     = "client_id"
	secret = "client_secret"
)

func OauthClientValidation() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (res interface{}, err error) {
		// start transaction
		sp := tracing.StartParent(grpchelper.SetTransactionName(ctx, info.FullMethod[strings.LastIndex(info.FullMethod, "/")+1:]+"Service"))
		defer tracing.Close(sp)

		// get metadata
		md, ok := metadata.FromIncomingContext(ctx)

		if !ok {
			return nil, tracing.LogError(sp, status.Error(response.ErrorFromCode(sp, 502)))
		}

		tracing.LogObject(sp, "metadata", md)

		// get oauth client id & secret
		if len(md[id]) == 0 && len(md[secret]) == 0 {
			return nil, tracing.LogError(sp, status.Error(response.ErrorFromCode(sp, 421)))
		}

		clientId := md[id][0]
		clientSecret := md[secret][0]

		// initiate oauth client use case
		oauthClientUC := container.InitializeOauthClientUsecase(repository.Con.MainMongoDB)

		// validate oauth client
		if err := oauthClientUC.ValidateOauthClient(tracing.Context(sp), struct {
			ClientId     string
			ClientSecret string
		}{
			ClientId:     clientId,
			ClientSecret: clientSecret,
		}); err != nil {
			return nil, tracing.LogError(sp, status.Error(response.Error(sp, err)))
		}

		return handler(tracing.Context(sp), req)
	}
}
