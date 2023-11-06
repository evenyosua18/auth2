package grpchelper

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

type GrpcHelper struct {
}

func (h *GrpcHelper) GetContextName(ctx interface{}) (context.Context, string) {
	//check nil
	if ctx == nil {
		return nil, ""
	}

	//to fiber
	c, ok := ctx.(context.Context)

	if !ok {
		return nil, ""
	}

	//check transaction name, use it if exist
	transactionName, ok := c.Value("transaction_name").(string)

	if !ok {
		transactionName = "Unknown Service"
	}

	return c, transactionName
}

func (h *GrpcHelper) GetInfo(ctx interface{}) (info map[string]interface{}) {
	//check nil
	if ctx == nil {
		return
	}

	//to fiber
	c, ok := ctx.(context.Context)

	if !ok {
		return
	}

	//convert to grpc metadata
	md, _ := metadata.FromIncomingContext(c)

	//set info
	info = make(map[string]interface{})
	info["authority"] = md["authority"]
	info["content-type"] = md["content-type"]
	info["grpc-accept-encoding"] = md["grpc-accept-encoding"]

	return
}
