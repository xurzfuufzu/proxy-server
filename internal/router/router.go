package router

import (
	"github.com/go-chi/chi/v5"
	"go-proxy/internal/service"
	"net/http"
)

func NewRouter() http.Handler {
	router := chi.NewRouter()

	router.Route("/v1", func(r chi.Router) {
		r.Get("/healthCheck", service.HealthCheck)
		r.Post("/proxy", service.DoProxy)
		r.Get("/proxy/{id}", service.GetByID)
	})

	return router
}
