package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Price struct {
	OID             primitive.ObjectID `bson:"_id" json:"-"`
	ID              string             `bson:"-" json:"id"`
	EntityID        string             `bson:"entity" json:"entity"`
	Source          string             `bson:"source" json:"source"`
	Fee             Fee                `bson:"fee" json:"fee"`
	Discount        Discount           `bson:"discount" json:"discount"`
	Commission      Commission         `bson:"commission" json:"commission"`
	AgentCommission AgentCommission    `bson:"agent_commission" json:"agent_commission"`
	TaxValue        float64            `bson:"-" json:"taxValue"`
	NetPrice        float64            `bson:"-" json:"netPrice"`
	NetPremium      float64            `bson:"-" json:"netPremium"`
}

func (p *Price) Calculate() {
	p.Discount.Calculate(p.NetPrice)
	p.Fee.Calculate(p.TaxValue)
	p.Commission.Calculate(p.NetPrice)
	p.AgentCommission.Calculate(p.NetPrice)
}
