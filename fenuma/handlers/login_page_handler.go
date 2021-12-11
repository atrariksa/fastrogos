package handlers

import (
	"html/template"
	"net/http"

	"github.com/atrariksa/fastrogos/fenuma/configs"
	"github.com/sirupsen/logrus"
)

const (
	LOGIN_PAGE_TITLE string = "Login"
	LOGIN_HTML       string = "login.html"
)

type LoginPage struct {
	Title string
}

type LoginPageHandler struct {
	cfg *configs.Config
	t   *template.Template
	log *logrus.Logger
}

func NewLoginPageHandler(cfg *configs.Config, t *template.Template, log *logrus.Logger) *LoginPageHandler {
	return &LoginPageHandler{
		cfg: cfg,
		t:   t,
		log: log,
	}
}

func (lp *LoginPageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	loginPage := LoginPage{Title: LOGIN_PAGE_TITLE}
	err := lp.t.ExecuteTemplate(w, LOGIN_HTML, loginPage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}
}
