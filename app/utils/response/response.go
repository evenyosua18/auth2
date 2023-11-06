package response

import "github.com/evenyosua18/ego-util/codes"

type DefaultResponse struct {
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	ErrorMessage string      `json:"errorMessage"`
	Data         interface{} `json:"data"`
}

func Set(data interface{}, code int) DefaultResponse {
	c := codes.Get(code)

	return DefaultResponse{
		Code:         code,
		Message:      c.ResponseMessage,
		ErrorMessage: c.ErrorMessage,
		Data:         data,
	}
}
