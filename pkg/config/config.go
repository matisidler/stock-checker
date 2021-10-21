package config

import "html/template"

type AppConfig struct {
	TplCache map[string]*template.Template
}
