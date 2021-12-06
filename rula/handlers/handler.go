package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/atrariksa/fastrogos/rula/configs"
	"github.com/atrariksa/fastrogos/rula/constants"
	"github.com/atrariksa/fastrogos/rula/docs"
	"github.com/atrariksa/fastrogos/rula/drivers"
	"github.com/atrariksa/fastrogos/rula/models"
	"github.com/atrariksa/fastrogos/rula/repos"
	"github.com/atrariksa/fastrogos/rula/services"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

type IHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	cfg *configs.Config
	log *logrus.Logger
	IReqValidator
	IRespWriter
}

func NewHandler(cfg *configs.Config, log *logrus.Logger) *Handler {
	return &Handler{
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
// @Router / [get]
func (lh *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	var resp = models.SuccessResp(http.StatusOK, "Hello")
	lh.Write(w, resp.HttpCode, resp)
	return
}

type IReqValidator interface {
	Validate(r *http.Request, req interface{}) (err error)
}

type ReqValidator struct {
}

func (rf *ReqValidator) Validate(r *http.Request, req interface{}) (err error) {
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(bodyByte, req)
	if err != nil {
		return
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		return
	}
	return
}

type IRespWriter interface {
	Write(w http.ResponseWriter, httpCode int, resp interface{})
}

type RespWriter struct {
}

func (rw *RespWriter) Write(w http.ResponseWriter, httpCode int, resp interface{}) {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(httpCode)
	rByte, _ := json.Marshal(&resp)
	w.Write(rByte)
}

func WireHandlers(r *chi.Mux, cfg *configs.Config, log *logrus.Logger) {

	setupSwagger(r, cfg, log)

	passwordHasher := &services.PasswordHasher{}
	passwordValidator := &services.PasswordValidator{}

	dbClient := drivers.NewDBClient(cfg)

	c := drivers.GetRedisClient(cfg)
	cacheRepo := repos.NewCache(cfg, c)
	userRepo := repos.UserRepo{DB: dbClient}
	userRepo.ICache = cacheRepo

	userService := services.NewUserService(cfg, log)
	userService.IUserRepo = &userRepo
	userService.PasswordHasher = passwordHasher
	userService.PasswordValidator = passwordValidator

	loginService := services.NewLoginService(&userService)

	reqValidator := ReqValidator{}
	respWriter := RespWriter{}

	generalHandler := NewHandler(cfg, log)
	generalHandler.IReqValidator = &reqValidator
	generalHandler.IRespWriter = &respWriter

	r.Get("/", generalHandler.Handle)

	r.Route(constants.USER_PATH, func(r chi.Router) {

		createUserHandler := NewCreateUserHandler(generalHandler)
		createUserHandler.SetService(&userService)

		updateUserHandler := NewUpdateUserHandler(generalHandler)
		updateUserHandler.SetService(&userService)

		deleteUserHandler := NewDeleteUserHandler(generalHandler)
		deleteUserHandler.SetService(&userService)

		r.Post("/", createUserHandler.Handle)
		r.Put("/", updateUserHandler.Handle)
		r.Delete("/", deleteUserHandler.Handle)

	})

	r.Route("/login", func(r chi.Router) {

		loginHandler := NewLoginHandler(generalHandler)
		loginHandler.SetService(&loginService)

		r.Post("/", loginHandler.Handle)
	})

}

func setupSwagger(r *chi.Mux, cfg *configs.Config, log *logrus.Logger) {
	docs.SwaggerInfo.Title = "Rula"
	docs.SwaggerInfo.Version = "0.0.1"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Host = cfg.App.Hostname
	swaggerURL := fmt.Sprintf("%s/swagger/doc.json", "http://"+cfg.App.Hostname)
	r.Get("/swagger/*", httpSwagger.Handler(httpSwagger.URL(swaggerURL)))
	log.Info("swagger url : ", swaggerURL)
}
