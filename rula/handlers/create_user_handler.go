package handlers

import (
	"net/http"

	"github.com/atrariksa/fastrogos/rula/models"
	"github.com/atrariksa/fastrogos/rula/services"
)

type CreateUserHandler struct {
	*Handler
	svc services.IUserService
}

func NewCreateUserHandler(h *Handler) *CreateUserHandler {
	return &CreateUserHandler{
		Handler: h,
	}
}

func (cuh *CreateUserHandler) SetService(svc services.IUserService) {
	cuh.svc = svc
}

// CreateUser example
// @Summary Create new User
// @Description Create new User
// @ID create-user
// @Accept  json
// @Produce  json
// @Param   models.CreateUserReq  body models.CreateUserReq true  "CreateUserReq"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /user/ [post]
func (cuh *CreateUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUserReq
	var resp models.Response
	err := cuh.Validate(r, &req)
	if err != nil {
		resp = models.ErrInvalidPayloadResp(err)
	} else {
		resp = cuh.svc.Create(req)
	}
	cuh.Write(w, resp.HttpCode, resp)
	return
}
