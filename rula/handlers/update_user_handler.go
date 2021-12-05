package handlers

import (
	"net/http"

	"github.com/atrariksa/fastrogos/rula/models"
	"github.com/atrariksa/fastrogos/rula/services"
)

type UpdateUserHandler struct {
	*Handler
	svc services.IUserService
}

func NewUpdateUserHandler(h *Handler) *UpdateUserHandler {
	return &UpdateUserHandler{
		Handler: h,
	}
}

func (uhh *UpdateUserHandler) SetService(svc services.IUserService) {
	uhh.svc = svc
}

// UpdateUser example
// @Summary Update User
// @Description Update User
// @ID Update User
// @Accept  json
// @Produce  json
// @Param   models.UpdateUserReq  body models.UpdateUserReq true  "UpdateUserReq"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /user/ [put]
func (uhh *UpdateUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req models.UpdateUserReq
	var resp models.Response
	err := uhh.Validate(r, &req)
	if err != nil {
		resp = models.ErrInvalidPayloadResp(err)
	} else {
		resp = uhh.svc.Update(req)
	}
	uhh.Handler.Write(w, resp.HttpCode, resp)
	return
}
