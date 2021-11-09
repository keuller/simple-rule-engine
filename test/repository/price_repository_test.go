package repository

import (
	"log"
	"testing"

	infra "github.com/keuller/pricing-service/internal/infrastructure/database"
)

func TestPriceRepositorySuite(t *testing.T) {
	if err := infra.NewConnection(); err != nil {
		t.FailNow()
	}

	repo := infra.NewPriceRepository()

	t.Run("Get Price By ID", func(it *testing.T) {
		price, err := repo.GetPrice("6185171247b155aff0db3663")
		if err != nil {
			t.FailNow()
		}

		log.Printf("[TEST] Price: %v \n", price)
	})

	t.Run("Get Price Using Invalid ID", func(it *testing.T) {
		_, err := repo.GetPrice("5f60faa9674119177ef00000")
		if err != nil {
			return
		}

		t.FailNow()
	})
}
