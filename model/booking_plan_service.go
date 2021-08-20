package model

import "time"

type BookingPlanService struct {
	ID            uint64    `gorm:"primaryKey;auto_increment" json:"id,omitempty"`
	UserID        uint64    `json:"user_id,omitempty"`
	PlanServiceID uint64    `json:"plan_service_id"`
	Status        string    `gorm:"default:request;not null" json:"status,omitempty"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

type BookingPlanServiceRepository interface {
	Create(*BookingPlanService) error
	Get(uint64) (*BookingPlanService, error)
	Update(*BookingPlanService) error
	CountUserBookPlanServiceByUserID(uint64) (uint64, error)
}
