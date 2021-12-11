package utils

import (
	"html/template"

	"github.com/atrariksa/fastrogos/fenuma/configs"
)

func GetHTMLTemplates(cfg *configs.Config) *template.Template {
	t := template.Must(template.ParseGlob(cfg.App.Page.TemplatesFolder + "/*.html"))
	return t
}
