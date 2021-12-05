package services

import (
	"net/http"

	"github.com/atrariksa/fastrogos/rula/errs"
	"github.com/atrariksa/fastrogos/rula/models"
)

type ILoginService interface {
	Login(req models.LoginReq) (resp models.Response)
}

type LoginService struct {
	*UserService
}

func NewLoginService(userService *UserService) LoginService {
	return LoginService{
		UserService: userService,
	}
}

func (ls *LoginService) Login(req models.LoginReq) (resp models.Response) {
	var user models.User
	user.Username = req.Username
	err := ls.GetUserByUsername(&user)
	if err != nil {
		ls.log.Error(err)
		switch err.Error() {
		case errs.ErrUserNotFound.Error():
			resp = models.ErrUserNotFoundResp(err)
		default:
			resp = models.ErrGeneralResp()
		}
		return
	}
	err = ls.Validate(user.Password, req.Password)
	if err != nil {
		resp = models.ErrUnauthorized(errs.ErrInvalidPassword)
		return
	}
	resp = models.SuccessResp(http.StatusOK, "Login success")
	return
}
