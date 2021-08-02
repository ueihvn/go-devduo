package model

import (
	"database/sql"
	"time"
)

type Profile struct {
	ID           uint           `json:"id,omitempty"`
	UserID       uint           `json:"user_id,omitempty"`
	Technologies []Technology   `gorm:"many2many:profile_technologies;" json:"technologies,omitempty"`
	Fields       []Field        `gorm:"many2many:profile_fields;" json:"fields,omitempty"`
	Contact      string         `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"contact,omitempty"`
	Description  sql.NullString `json:"description,omitempty"`
	CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

type ProfileRepository interface {
	Create(interface{}) error
	// Get(int) (*Profile, error)
	// Update(*Profile) error
	// Delete(*Profile) error
}
