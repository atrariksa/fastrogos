package services

import (
	"github.com/atrariksa/fastrogos/corey/configs"
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

func (las *LoginAPIService) Login(req LoginAPIRequest) {

}
