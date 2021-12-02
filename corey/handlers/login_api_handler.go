package handlers

import (
	"net/http"

	"github.com/atrariksa/fastrogos/corey/configs"
	"github.com/atrariksa/fastrogos/corey/handlers/services"
	"github.com/sirupsen/logrus"
)

type LoginAPIHandler struct {
	cfg      *configs.Config
	log      *logrus.Logger
	loginsvc *services.LoginAPIService
}

func NewLoginAPIHandler(cfg *configs.Config, log *logrus.Logger) *LoginAPIHandler {
	return &LoginAPIHandler{
		cfg: cfg,
		log: log,
	}
}

func (lah *LoginAPIHandler) Handle(w http.ResponseWriter, r *http.Request) {

}
