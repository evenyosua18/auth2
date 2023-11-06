package grpchelper

import "context"

func SetTransactionName(ctx context.Context, name string) context.Context {
	return context.WithValue(ctx, "transaction_name", name)
}
