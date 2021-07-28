package model

import "time"

type BookingPlanService struct {
	ID            uint `gorm:"primary_key;auto_increment"`
	UserID        uint `gorm:"primaryKey"`
	PlanServiceID uint `gorm:"primaryKey"`
	Status        string
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
