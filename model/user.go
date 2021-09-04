package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                  uint64               `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Email               string               `gorm:"not null;unique" json:"email,omitempty"`
	Password            string               `gorm:"not null" json:"password,omitempty"`
	Profile             Profile              `json:"-"`
	PlanServices        []PlanService        `json:"-"`
	BookingPlanServices []BookingPlanService `json:"-"`
	CreatedAt           time.Time            `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt           time.Time            `gorm:"default:CURRENT_TIMESTAMP" json:"-"`
}

type UserRepository interface {
	CreateUser(*User) error
	GetUserById(uint64) (*User, error)
	GetUserByEmail(string) (*User, error)
	GetAllUsers() ([]User, error)
	UpdateUser(*User) error
	InitData() error
}

func (user *User) HashPassword() (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil

}

func (user *User) ComparePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}
	return nil
}
