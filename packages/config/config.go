package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

type AppConfig struct {
	UseCache     bool
	InProduction bool

	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger

	Session *scs.SessionManager
}
