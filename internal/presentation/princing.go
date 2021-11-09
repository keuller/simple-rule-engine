package presentation

import (
	"log"
	"net/http"

	"github.com/keuller/pricing-service/internal/domain/repository"
	domain "github.com/keuller/pricing-service/internal/domain/service"
	infra "github.com/keuller/pricing-service/internal/infrastructure/database"
)

var (
	repo        repository.PriceRepository
	service     domain.PriceService
	initialized bool
)

func initialize() {
	if initialized {
		return
	}

	log.Println("Loading runtime....")
	repo = infra.NewPriceRepository()
	service = domain.NewPriceService(repo)
	initialized = true
}

func Index() http.HandlerFunc {
	initialize()
	return func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Header().Set("Content-Type", "text/plain")
		res.Write([]byte("Princing service is working."))
	}
}

func GetPlans() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		res.WriteHeader(200)
		res.Header().Set("Content-Type", "text/plain")
		res.Write([]byte("Hello princing service"))
	}
}

func GetPrices() http.HandlerFunc {
	initialize()
	return func(res http.ResponseWriter, req *http.Request) {
		prices := service.GetPrices()
		Json(res, prices)
	}
}

func GetPrice() http.HandlerFunc {
	initialize()
	return func(res http.ResponseWriter, req *http.Request) {
		var data domain.PriceRequest
		FromJson(req.Body, &data)
		price, err := service.GetPlanPrice(data)
		if err != nil {
			res.WriteHeader(400)
			res.Header().Set("Content-Type", "application/json")
			res.Write([]byte(err.Error()))
			return
		}

		Json(res, price)
	}
}

func ReloadRules() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		service.ReloadRules()
		Json(res, map[string]string{
			"message": "Rules reload successfuly.",
		})
	}
}
