package handlers

import (
	"net/http"

	"github.com/atrariksa/fastrogos/rula/models"
	"github.com/atrariksa/fastrogos/rula/services"
)

type LoginHandler struct {
	*Handler
	svc services.ILoginService
}

func NewLoginHandler(h *Handler) *LoginHandler {
	return &LoginHandler{
		Handler: h,
	}
}

func (lh *LoginHandler) SetService(svc services.ILoginService) {
	lh.svc = svc
}

// Login example
// @Summary Login
// @Description Login
// @ID Login
// @Accept  json
// @Produce  json
// @Param   models.LoginReq  body models.LoginReq true  "LoginReq"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 401 {object} models.Response
// @Router /login/ [post]
func (lh *LoginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req models.LoginReq
	var resp models.Response
	err := lh.Validate(r, &req)
	if err != nil {
		resp = models.ErrInvalidPayloadResp(err)
	} else {
		resp = lh.svc.Login(req)
	}
	lh.Write(w, resp.HttpCode, resp)
	return
}
