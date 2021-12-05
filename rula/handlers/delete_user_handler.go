package handlers

import (
	"net/http"

	"github.com/atrariksa/fastrogos/rula/models"
	"github.com/atrariksa/fastrogos/rula/services"
)

type DeleteUserHandler struct {
	*Handler
	svc services.IUserService
}

func NewDeleteUserHandler(h *Handler) *DeleteUserHandler {
	return &DeleteUserHandler{
		Handler: h,
	}
}

func (duh *DeleteUserHandler) SetService(svc services.IUserService) {
	duh.svc = svc
}

// DeleteUser example
// @Summary DeleteUser
// @Description DeleteUser
// @ID Delete User
// @Accept  json
// @Produce  json
// @Param   models.DeleteUserReq  body models.DeleteUserReq true  "DeleteUserReq"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /user/ [delete]
func (duh *DeleteUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req models.DeleteUserReq
	var resp models.Response
	err := duh.Validate(r, &req)
	if err != nil {
		resp = models.ErrInvalidPayloadResp(err)
	} else {
		resp = duh.svc.Delete(req)
	}
	duh.Write(w, resp.HttpCode, resp)
	return
}
