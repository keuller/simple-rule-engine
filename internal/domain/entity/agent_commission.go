package entity

import "github.com/keuller/pricing-service/internal/common"

type AgentCommission struct {
	Type      string  `bson:"type" json:"type"`
	Condition string  `bson:"value" json:"condition"`
	Value     float64 `bson:"-" json:"value"`
}

func (a AgentCommission) IsValid() bool {
	return true
}

func (a *AgentCommission) Calculate(netValue float64) {
	val := common.GetValue(a.Type, a.Condition)
	if a.Type == "percentage" {
		a.Value = netValue * val
		return
	}
	a.Value = val
}
