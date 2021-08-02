package model

import "time"

type User struct {
	ID                  uint          `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	FullName            string        `gorm:"not null" json:"full_name,omitempty"`
	Email               string        `gorm:"not null;unique" json:"email,omitempty"`
	UserName            string        `gorm:"not null;unique" json:"user_name,omitempty"`
	Password            string        `gorm:"not null" json:"password,omitempty"`
	Profile             Profile       `json:"profile,omitempty"`
	PlanServices        []PlanService `json:"plan_services,omitempty"`
	BookingPlanServices []PlanService `gorm:"many2many:booking_plan_service;" json:"booking_plan_services,omitempty"`
	CreatedAt           time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt           time.Time     `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

type UserRepository interface {
	CreateUser(*User) error
	GetUserById(uint) (*User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(*User) error
}
