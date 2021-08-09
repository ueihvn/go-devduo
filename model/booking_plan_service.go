package model

import "time"

type BookingPlanService struct {
	ID            uint64 `gorm:"primaryKey;auto_increment"`
	UserID        uint64
	PlanServiceID uint64
	Status        string    `gorm:"default:request;not null"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}

type BookingPlanServiceRepository interface {
	Create(*BookingPlanService) error
	Get(uint) (*BookingPlanService, error)
	Update(*BookingPlanService) error
}
