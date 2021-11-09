package entity

import "github.com/keuller/pricing-service/internal/common"

type Fee struct {
	Type      string  `bson:"type" json:"type"`
	Condition string  `bson:"value" json:"condition"`
	Value     float64 `bson:"-" json:"value"`
}

func (f Fee) IsValid() bool {
	return true
}

func (f *Fee) Calculate(taxValue float64) {
	val := common.GetValue(f.Type, f.Condition)
	if f.Type == "percentage" {
		f.Value = taxValue * val
		return
	}
	f.Value = val
}
