package services

import (
	"context"
	"net/http"

	"github.com/atrariksa/fastrogos/fenuma/configs"
	"github.com/atrariksa/fastrogos/fenuma/handlers/client_services"
	"github.com/atrariksa/fastrogos/fenuma/models"
	rulaOp "github.com/atrariksa/fastrogos/rula/client/operations"
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
		if v, ok := err.(*rulaOp.LoginUnauthorized); ok {
			return models.Response{
				HttpCode: http.StatusUnauthorized,
				Code:     v.Payload.Code,
				Message:  v.Payload.Message,
			}
		}
		if v, ok := err.(*rulaOp.LoginBadRequest); ok {
			return models.Response{
				HttpCode: http.StatusBadRequest,
				Code:     v.Payload.Code,
				Message:  v.Payload.Message,
			}
		}
		return models.ErrGeneralResp()
	}
	res := models.SuccessResp(http.StatusOK, resp.Payload.Message)
	res.Data = resp.Payload.Data
	return res
}
