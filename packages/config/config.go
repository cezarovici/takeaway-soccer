package config

import (
	"html/template"
	"log"
)

type AppConfig struct {
	// todo add template cache
	UseCache      bool
	TemplateCache map[string]*template.Template

	InfoLog *log.Logger
}
