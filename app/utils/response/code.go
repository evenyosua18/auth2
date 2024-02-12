package response

import (
	code "github.com/evenyosua18/codes"
	"github.com/evenyosua18/tracing"
	"google.golang.org/grpc/codes"
)

func Code(customCode int) codes.Code {
	return codes.Code(customCode)
}

func Error(span interface{}, err error) (codes.Code, string) {
	c := code.Extract(err)

	tracing.LogObject(span, "codes", c)

	return codes.Code(c.ResponseCode), c.ResponseMessage
}

func ErrorFromCode(span interface{}, codeNumber int) (codes.Code, string) {
	c := code.Get(codeNumber)

	tracing.LogObject(span, "codes", c)

	return codes.Code(c.ResponseCode), c.ResponseMessage
}
