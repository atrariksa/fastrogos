package services

import (
	"context"
	"net/http"

	"github.com/atrariksa/fastrogos/fenuma/configs"
	"github.com/atrariksa/fastrogos/fenuma/handlers/client_services"
	"github.com/atrariksa/fastrogos/fenuma/models"
	"github.com/sirupsen/logrus"
)

type LoginAPIResponse struct {
}

type LoginAPIService struct {
	cfg *configs.Config
	log *logrus.Logger
}

func GetLoginAPIService(cfg *configs.Config, log *logrus.Logger) *LoginAPIService {
	return &LoginAPIService{
		cfg: cfg,
		log: log,
	}
}

func (las *LoginAPIService) Process(ctx context.Context, req interface{}) interface{} {
	loginReq := req.(models.LoginReq)
	resp, err := client_services.Login(loginReq)
	if err != nil {
		return models.ErrGeneralResp()
	}
	res := models.SuccessResp(http.StatusOK, resp.Payload.Message)
	res.Data = resp.Payload.Data
	return res
}
