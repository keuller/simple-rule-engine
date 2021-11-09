package service

import (
	"log"

	"github.com/keuller/pricing-service/internal/domain/entity"
	"github.com/keuller/pricing-service/internal/domain/fact"
	"github.com/keuller/pricing-service/internal/domain/repository"
)

type PriceService struct {
	ruleBase PriceRuleBase
	repo     repository.PriceRepository
}

func NewPriceService(repo repository.PriceRepository) PriceService {
	rb := NewPriceRuleBase(repo)
	return PriceService{rb, repo}
}

func (ps PriceService) GetPrices() []entity.Price {
	prices := ps.repo.GetPrices()
	if prices == nil {
		return []entity.Price{}
	}
	return prices
}

func (ps PriceService) GetPlanPrice(data PriceRequest) (entity.Price, error) {
	price, err := ps.repo.GetPrice(data.PlanId)
	if err != nil {
		return entity.Price{}, err
	}

	fact := &fact.DeviceProtectionFact{
		DeviceValue: data.DeviceValue,
		NetPrice:    0.0,
		TaxValue:    0.0,
	}

	ps.ruleBase.AddFact("DP", fact)
	ps.ruleBase.AddFact("Price", &price)
	ps.ruleBase.Process()
	price.ID = data.PlanId
	return price, nil
}

func (ps PriceService) ReloadRules() {
	log.Println("Reloading rules base...")
	ps.ruleBase.BuildRuleBase()
}
