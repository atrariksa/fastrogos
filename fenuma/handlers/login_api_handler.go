package handlers

import (
	"net/http"

	api_utils "github.com/atrariksa/api_utils"
	"github.com/atrariksa/fastrogos/fenuma/configs"
	"github.com/atrariksa/fastrogos/fenuma/models"
	"github.com/sirupsen/logrus"
)

type LoginAPIHandler struct {
	cfg *configs.Config
	log *logrus.Logger
	api_utils.DefaultHttpHandler
}

func NewLoginAPIHandler(cfg *configs.Config, log *logrus.Logger) *LoginAPIHandler {
	return &LoginAPIHandler{
		cfg: cfg,
		log: log,
	}
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
// @Router /api/login [post]
func (lah *LoginAPIHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req models.LoginReq
	var resp interface{}
	err := lah.Unmarshal(r, &req)
	if err != nil {
		lah.Write(w, models.ErrGeneralResp().HttpCode, models.ErrGeneralResp())
		return
	}
	resp = lah.Process(r.Context(), req)
	lah.Write(w, resp.(models.Response).HttpCode, resp)
	return
}
