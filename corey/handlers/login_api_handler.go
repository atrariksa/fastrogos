package handlers

import (
	"net/http"

	"github.com/atrariksa/gocosmos/corey/configs"
	"github.com/atrariksa/gocosmos/corey/handlers/services"
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
