package service

type PriceRequest struct {
	PlanId      string  `json:"planId"`
	DeviceValue float64 `json:"device"`
}
