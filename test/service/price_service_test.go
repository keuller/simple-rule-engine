package service

import (
	"log"
	"testing"

	domain "github.com/keuller/pricing-service/internal/domain/service"
	infra "github.com/keuller/pricing-service/internal/infrastructure/database"
)

func TestPriceServiceSuite(t *testing.T) {
	repo := infra.NewPriceRepository()
	service := domain.NewPriceService(repo)

	t.Run("Get Plan Price", func(it *testing.T) {
		data := domain.PriceRequest{"5f60faa9674119177eff18ab", 12000.5}
		price, err := service.GetPlanPrice(data)
		if err != nil {
			t.FailNow()
		}

		log.Printf("[Test] DeviceValue: %.2f \n", data.DeviceValue)
		log.Printf("[Test] NetPrice...: %.2f \n", price.NetPrice)
		log.Printf("[Test] TaxValue...: %.2f \n", price.TaxValue)
		log.Printf("[Test] NetPremium.: %.2f \n", price.NetPremium)
	})
}
