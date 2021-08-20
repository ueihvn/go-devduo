package model

import "time"

type User struct {
	ID                  uint64               `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Email               string               `gorm:"not null;unique" json:"email,omitempty"`
	Password            string               `gorm:"not null" json:"password,omitempty"`
	Profile             Profile              `json:"-"`
	PlanServices        []PlanService        `json:"plan_services,omitempty"`
	BookingPlanServices []BookingPlanService `json:"booking_plan_services,omitempty"`
	CreatedAt           time.Time            `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt           time.Time            `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}

type UserRepository interface {
	CreateUser(*User) error
	GetUserById(uint64) (*User, error)
	GetUserByEmail(string) (*User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(*User) error
}
