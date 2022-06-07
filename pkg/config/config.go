package config

import (
	"text/template"

	"github.com/alexedwards/scs/v2"
)

//App config holds the application config
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	Session       *scs.SessionManager
}
