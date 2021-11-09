package entity

import "github.com/keuller/pricing-service/internal/common"

type Discount struct {
	Type      string  `bson:"type" json:"type"`
	Condition string  `bson:"value" json:"condition"`
	Value     float64 `bson:"-" json:"value"`
}

func (d Discount) IsValid() bool {
	return true
}

func (d *Discount) Calculate(netValue float64) {
	val := common.GetValue(d.Type, d.Condition)
	if d.Type == "percentage" {
		d.Value = netValue * val
		return
	}
	d.Value = val
}
