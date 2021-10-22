package handlers

import (
	"net/http"

	"github.com/matisidler/stock-checker/pkg/config"
	"github.com/matisidler/stock-checker/pkg/render"
)

var Repo *Repository

type Repository struct {
	Config config.AppConfig
}

func NewRepo(config config.AppConfig) *Repository {
	return &Repository{
		Config: config,
	}
}

func (m Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.ExecutingTemplate(w, r, "home.page.tpl")
}
