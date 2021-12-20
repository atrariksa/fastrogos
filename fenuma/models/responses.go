package models

import "net/http"

type Response struct {
	HttpCode int         `json:"-"`
	Code     string      `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data"`
}

func SuccessResp(httpCode int, message string) Response {
	return Response{
		HttpCode: httpCode,
		Code:     "00000",
		Message:  message,
	}
}

func ErrGeneralResp() Response {
	return Response{
		HttpCode: http.StatusInternalServerError,
		Code:     "err-ups-10001",
		Message:  "Oops, something went wrong",
	}
}
