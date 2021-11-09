package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/keuller/pricing-service/internal/domain/entity"
	domain "github.com/keuller/pricing-service/internal/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	timeout = 5 * time.Second
)

type priceRepository struct {
}

func NewPriceRepository() domain.PriceRepository {
	return priceRepository{}
}

func (pr priceRepository) GetPrices() []entity.Price {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	collection := db.Collection("prices")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("[WARN] fail to fetch data - %s \n", err.Error())
		return nil
	}

	prices := make([]entity.Price, 0)
	var result entity.Price
	for cursor.Next(ctx) {
		_ = cursor.Decode(&result)
		result.ID = result.OID.Hex()
		prices = append(prices, result)
	}

	cursor.Close(ctx)
	return prices
}

func (pr priceRepository) GetPrice(id string) (entity.Price, error) {
	var result entity.Price

	pr.GetPrices()

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	collection := db.Collection("prices")
	oid, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{"_id": bson.M{"$eq": oid}}

	if err := collection.FindOne(ctx, condition).Decode(&result); err != nil {
		log.Printf("[WARN] %s \n", err.Error())
		return entity.Price{}, fmt.Errorf("price not found")
	}

	return result, nil
}

func (pr priceRepository) GetRules() []entity.Rule {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	collection := db.Collection("rules")
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("[WARN] fail to fetch data - %s \n", err.Error())
		return nil
	}

	rules := make([]entity.Rule, 0)
	var result entity.Rule
	for cursor.Next(ctx) {
		_ = cursor.Decode(&result)
		rules = append(rules, result)
	}

	cursor.Close(ctx)
	return rules
}
