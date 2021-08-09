package model

import "github.com/shopspring/decimal"

type PlanService struct {
	ID                  uint64 `gorm:"auto_increment"`
	UserID              uint64
	Title               string          `gorm:"unique;not null"`
	Description         string          `gorm:"null"`
	Price               decimal.Decimal `gorm:"type:decimal(10,2);not null"`
	BookingPlanServices []BookingPlanService
}

type PlanServiceRepository interface {
	Create(*PlanService) error
	Get(uint64) (*PlanService, error)
	Update(*PlanService) error
}
