package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/keuller/pricing-service/internal/presentation"
)

func registerRoutes(router *chi.Mux) {
	router.Get("/", presentation.Index())
	router.Get("/v1/plans", presentation.GetPlans())
	router.Get("/v1/prices", presentation.GetPrices())
	router.Post("/v1/price", presentation.GetPrice())
	router.Get("/v1/prices/reload", presentation.ReloadRules())
}
