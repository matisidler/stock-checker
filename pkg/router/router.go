package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/matisidler/stock-checker/pkg/config"
	"github.com/matisidler/stock-checker/pkg/handlers"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)
	return mux
}
