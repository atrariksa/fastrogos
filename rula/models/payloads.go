package models

import "net/http"

type CreateUserReq struct {
	Username string `json:"username" valid:"required~Username cannot not be empty"`
	Email    string `json:"email" valid:"email~Please fill valid email"`
	Address  string `json:"address" valid:"required~Address cannot be empty"`
	Password string `json:"password" valid:"required~Password cannot be empty"`
}

type UpdateUserReq struct {
	ID       uint32 `json:"id" valid:"required~ID cannot not be empty"`
	UserID   string `json:"user_id" valid:"required~Username cannot not be empty"`
	Username string `json:"username" valid:"required~Username cannot not be empty"`
	Email    string `json:"email" valid:"email~Please fill valid email"`
	Address  string `json:"address" valid:"required~Address cannot be empty"`
	Password string `json:"password" valid:"required~Password cannot be empty"`
}

type DeleteUserReq struct {
	Username string `json:"username" valid:"required~Username cannot not be empty"`
}

type LoginReq struct {
	Username string `json:"username" valid:"required~Username cannot be empty"`
	Password string `json:"password" valid:"required~Password cannot be empty"`
}

type Response struct {
	HttpCode int         `json:"-"`
	Code     string      `json:"code"`
	Message  string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
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

func ErrInvalidPayloadResp(err error) Response {
	return Response{
		HttpCode: http.StatusBadRequest,
		Code:     "err-ups-10002",
		Message:  err.Error(),
	}
}

func ErrUserNotFoundResp(err error) Response {
	return Response{
		HttpCode: http.StatusNotFound,
		Code:     "err-ups-10003",
		Message:  err.Error(),
	}
}

func ErrDuplicateKeyResp(err error) Response {
	return Response{
		HttpCode: http.StatusConflict,
		Code:     "err-ups-10004",
		Message:  err.Error(),
	}
}

func ErrUnauthorized(err error) Response {
	return Response{
		HttpCode: http.StatusUnauthorized,
		Code:     "err-ups-10005",
		Message:  err.Error(),
	}
}
