package middlewares

import (
	"net/http"
	"strings"

	"github.com/atrariksa/fastrogos/corey/configs"
)

func DefaultResponseHeadersHandler(cfg *configs.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return responseHeaders(next, cfg)
	}
}

func responseHeaders(next http.Handler, cfg *configs.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for k, v := range getDefaultRespHeaders(cfg) {
			w.Header().Add(k, v)
		}
		next.ServeHTTP(w, r)
	})
}

func getDefaultRespHeaders(cfg *configs.Config) map[string]string {
	headerMap := map[string]string{}
	headerMap["X-Frame-Options"] = cfg.App.DefaultRespHeaders.XFrameOptions
	headerMap["X-Content-Type-Options"] = cfg.App.DefaultRespHeaders.XContentTypeOptions
	headerMap["X-XSS-Protection"] = cfg.App.DefaultRespHeaders.XXSSProtection
	headerMap["Content-Security-Policy"] = strings.Join(cfg.App.DefaultRespHeaders.CSP, ";")
	return headerMap
}
