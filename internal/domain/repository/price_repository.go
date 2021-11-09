package repository

import "github.com/keuller/pricing-service/internal/domain/entity"

type PriceRepository interface {
	GetPrice(id string) (entity.Price, error)
	GetPrices() []entity.Price
	GetRules() []entity.Rule
}
