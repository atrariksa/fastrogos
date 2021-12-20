package handlers

import (
	"net/http"

	api_utils "github.com/atrariksa/api_utils"
	"github.com/atrariksa/fastrogos/fenuma/configs"
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

// example
// @Summary
// @Description
// @ID general
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/login [post]
func (lah *LoginAPIHandler) Handle(w http.ResponseWriter, r *http.Request) {

}
