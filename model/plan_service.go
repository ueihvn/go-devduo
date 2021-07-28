package model

import "github.com/shopspring/decimal"

type PlanService struct {
	ID     uint `gorm:"auto_increment"`
	UserID uint
	Title  string          `gorm:"not null"`
	Price  decimal.Decimal `gorm:"type:decimal(10,2);not null"`
}
