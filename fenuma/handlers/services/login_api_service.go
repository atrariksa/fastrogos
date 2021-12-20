package services

import (
	"context"
	"net/http"

	"github.com/atrariksa/fastrogos/fenuma/configs"
	"github.com/atrariksa/fastrogos/fenuma/models"
	"github.com/sirupsen/logrus"
)

type LoginAPIRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

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
	return models.SuccessResp(http.StatusOK, "Success")
}
