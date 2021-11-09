package entity

import "github.com/keuller/pricing-service/internal/common"

type Commission struct {
	Type      string  `bson:"type" json:"type"`
	Condition string  `bson:"value" json:"condition"`
	Value     float64 `bson:"-" json:"value"`
}

func (c Commission) IsValid() bool {
	return true
}

func (c *Commission) Calculate(netValue float64) {
	val := common.GetValue(c.Type, c.Condition)
	if c.Type == "percentage" {
		c.Value = netValue * val
		return
	}
	c.Value = val
}
