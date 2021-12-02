package utils

import (
	"github.com/atrariksa/fastrogos/corey/configs"
	"github.com/go-chi/cors"
)

func GetCorsOptions(cfg *configs.Config) cors.Options {
	corsCfg := cfg.App.CORS
	c := cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: corsCfg.AllowedOrigins,
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods: corsCfg.AllowedMethods,
		AllowedHeaders: corsCfg.AllowedHeaders,
		// ExposedHeaders:   []string{"Link"},
		AllowCredentials: corsCfg.AllowCredentials,
		MaxAge:           corsCfg.MaxAgeSeconds, // Maximum value not ignored by any of major browsers
	}
	return c
}
