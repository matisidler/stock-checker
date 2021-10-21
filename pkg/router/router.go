package router

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/matisidler/stock-checker/pkg/handlers"
)

func Routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Home)
	return mux
}
