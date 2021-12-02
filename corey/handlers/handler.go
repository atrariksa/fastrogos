package handlers

import (
	"net/http"

	"github.com/atrariksa/fastrogos/corey/configs"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/tools/go/packages"
)

type IHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type ChiMidleware struct {
	Packages *[]packages.Package
}

func WireHandlers(r *chi.Mux, cfg *configs.Config, log *logrus.Logger) {

	r.Route("/api", func(r chi.Router) {
		r.Post("/login", func(rw http.ResponseWriter, r *http.Request) {

		})
	})
}
