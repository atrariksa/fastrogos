package services

import (
	"net/http"
	"strings"

	"github.com/atrariksa/fastrogos/rula/configs"
	"github.com/atrariksa/fastrogos/rula/errs"
	"github.com/atrariksa/fastrogos/rula/models"
	"github.com/atrariksa/fastrogos/rula/repos"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	Create(req models.CreateUserReq) (resp models.Response)
	Update(req models.UpdateUserReq) (resp models.Response)
	Delete(req models.DeleteUserReq) (resp models.Response)
}

type UserService struct {
	cfg *configs.Config
	log *logrus.Logger
	repos.IUserRepo
	*PasswordHasher
	*PasswordValidator
}

func NewUserService(cfg *configs.Config, log *logrus.Logger) UserService {
	return UserService{
		cfg: cfg,
		log: log,
	}
}

func (us *UserService) Create(req models.CreateUserReq) (resp models.Response) {
	newUser := models.User{
		Username: req.Username,
		Email:    req.Email,
	}
	hashedPassword, err := us.Hash(req.Password)
	if err != nil {
		us.log.Error(err)
		resp = models.ErrGeneralResp()
		return
	}
	newUser.Password = hashedPassword
	err = us.CreateUser(newUser)
	if err != nil {
		us.log.Error(err)
		switch {
		case strings.Contains(err.Error(), repos.DuplicateKey):
			resp = models.ErrDuplicateKeyResp(err)
		default:
			resp = models.ErrGeneralResp()
		}
		return
	}
	resp = models.SuccessResp(http.StatusCreated, "User created")
	return
}

func (us *UserService) Update(req models.UpdateUserReq) (resp models.Response) {
	updateData := models.User{
		ID:       uint(req.ID),
		Username: req.Username,
		Email:    req.Email,
	}
	hashedPassword, err := us.Hash(req.Password)
	if err != nil {
		us.log.Error(err)
		resp = models.ErrGeneralResp()
		return
	}
	updateData.Password = hashedPassword
	err = us.UpdateUser(updateData)
	if err != nil {
		us.log.Error(err)
		switch err.Error() {
		case errs.ErrUserNotFound.Error():
			resp = models.ErrUserNotFoundResp(err)
		default:
			resp = models.ErrGeneralResp()
		}
		return
	}
	resp = models.SuccessResp(http.StatusOK, "User updated")
	return
}

func (us *UserService) Delete(req models.DeleteUserReq) (resp models.Response) {
	deleteData := models.User{
		Username: req.Username,
	}
	err := us.GetUserByUsername(&deleteData)
	if err != nil {
		us.log.Error(err)
		switch err.Error() {
		case errs.ErrUserNotFound.Error():
			resp = models.ErrUserNotFoundResp(err)
		default:
			resp = models.ErrGeneralResp()
		}
		return
	}
	err = us.DeleteUser(deleteData)
	if err != nil {
		us.log.Error(err)
		switch err.Error() {
		case errs.ErrUserNotFound.Error():
			resp = models.ErrUserNotFoundResp(err)
		default:
			resp = models.ErrGeneralResp()
		}
		return
	}
	resp = models.SuccessResp(http.StatusOK, "User deleted")
	return
}

type IPasswordValidator interface {
	Validate(hashedPassword, password string) error
}
type PasswordValidator struct{}

func (pv *PasswordValidator) Validate(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

type IPasswordHasher interface {
	Hash(password string) (string, error)
}
type PasswordHasher struct{}

func (pv *PasswordHasher) Hash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
