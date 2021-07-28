package model

import (
	"database/sql"
	"time"
)

type Profile struct {
	ID                   uint
	UserID               uint
	ProgrammingLanguages []ProgrammingLanguage `gorm:"many2many:profile_programming_languages;"`
	Fields               []Field               `gorm:"many2many:profile_fields;"`
	Contact              string                `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB"`
	Description          sql.NullString
	CreatedAt            time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt            time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
