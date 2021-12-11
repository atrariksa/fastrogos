package handlers

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/atrariksa/fastrogos/fenuma/configs"
	"github.com/go-chi/chi/v5"
)

type StaticHandler struct {
	cfg *configs.Config
}

func NewStaticHandler(cfg *configs.Config) *StaticHandler {
	return &StaticHandler{
		cfg: cfg,
	}
}

func (sh *StaticHandler) Handle(w http.ResponseWriter, r *http.Request) {
	workDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	root := http.Dir(filepath.Join(workDir, sh.cfg.App.Page.AssetsFolder))
	rctx := chi.RouteContext(r.Context())
	pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
	fs := http.StripPrefix(pathPrefix, http.FileServer(root))
	fs.ServeHTTP(w, r)
}
