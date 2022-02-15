package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

// AppConfig will hold the application config
// This allows sitewide access to anyting in this file
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InProduction  bool
	InfoLog       *log.Logger
	Session       *scs.SessionManager
}
