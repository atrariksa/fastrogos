package utils

import (
	"github.com/atrariksa/fastrogos/rula/configs"
	"github.com/go-chi/cors"
)

func GetCorsOptions(cfg *configs.Config) cors.Options {
	corsCfg := cfg.App.CORS
	c := cors.Options{
		AllowedOrigins:   corsCfg.AllowedOrigins,
		AllowedMethods:   corsCfg.AllowedMethods,
		AllowedHeaders:   corsCfg.AllowedHeaders,
		AllowCredentials: corsCfg.AllowCredentials,
		MaxAge:           corsCfg.MaxAgeSeconds,
	}
	return c
}
