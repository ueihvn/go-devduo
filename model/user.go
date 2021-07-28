package model

import "time"

type User struct {
	ID                  uint   `gorm:"primary_key;auto_increment"`
	FullName            string `gorm:"not null"`
	Email               string `gorm:"not null;unique"`
	UserName            string `gorm:"not null;unique"`
	Password            string `gorm:"not null"`
	Profile             Profile
	PlanServices        []PlanService
	BookingPlanServices []PlanService `gorm:"many2many:booking_plan_service;"`
	CreatedAt           time.Time     `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt           time.Time     `gorm:"default:CURRENT_TIMESTAMP"`
}
