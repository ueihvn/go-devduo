package model

import "github.com/shopspring/decimal"

type PlanService struct {
	ID                  uint64               `gorm:"auto_increment" json:"id,omitempty"`
	UserID              uint64               `json:"user_id,omitempty"`
	Title               string               `gorm:"unique;not null" json:"title,omitempty"`
	Description         string               `gorm:"null" json:"description,omitempty"`
	Price               decimal.Decimal      `gorm:"type:decimal(10,2);not null" json:"price,omitempty"`
	BookingPlanServices []BookingPlanService `json:"booking_plan_services,omitempty"`
}

type PlanServiceRepository interface {
	Create(*PlanService) error
	Get(uint64) (*PlanService, error)
	Update(*PlanService) error
}
