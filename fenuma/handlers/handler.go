package handlers

import (
	"net/http"

	apiutils "github.com/atrariksa/api_utils"
	"github.com/atrariksa/fastrogos/fenuma/configs"
	"github.com/atrariksa/fastrogos/fenuma/handlers/services"
	"github.com/atrariksa/fastrogos/fenuma/utils"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/tools/go/packages"
)

type ChiMidleware struct {
	Packages *[]packages.Package
}

func WireHandlers(r *chi.Mux, cfg *configs.Config, log *logrus.Logger) {

	staticHandler := NewStaticHandler(cfg)
	r.Get("/assets/*", staticHandler.Handle)
	r.Get("/favicon.ico", http.RedirectHandler("/assets/favicon.ico", 301).ServeHTTP)

	htmlTemplates := utils.GetHTMLTemplates(cfg)
	loginPageHandler := NewLoginPageHandler(cfg, htmlTemplates, log)
	r.Get("/login", loginPageHandler.Handle)

	dh := apiutils.GetDefaultHandler()
	loginAPIHandler := NewLoginAPIHandler(cfg, log)
	loginAPIHandler.DefaultHttpHandler = dh
	loginAPIHandler.IDefaultService = services.GetLoginAPIService(cfg, log)

	r.Route("/api", func(r chi.Router) {
		r.Post("/login", loginAPIHandler.Handle)
	})
}
