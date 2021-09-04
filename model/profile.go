package model

import (
	"database/sql"
	"encoding/json"
	"time"
)

type Profile struct {
	UserID       uint64         `gorm:"primaryKey" json:"user_id,omitempty"`
	FullName     string         `gorm:"not null" json:"full_name,omitempty"`
	Technologies []Technology   `gorm:"many2many:profile_technologies;foreignKey:UserID;joinForeignkey:ProfileUserID" json:"technologies,omitempty"`
	Fields       []Field        `gorm:"many2many:profile_fields;foreignKey:UserID;joinForeignkey:ProfileUserID" json:"fields,omitempty"`
	Contact      string         `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"contact,omitempty"`
	ThumnailURL  string         `gorm:"null" json:"thumnail_url,omitempty"`
	Description  sql.NullString `json:"description,omitempty"`
	CreatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"created_at,omitempty"`
	UpdatedAt    time.Time      `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at,omitempty"`
}

// swagger:model Profile
type ProfileJSON struct {
	// id for user. required for Create, Update, Get
	// required: true
	// min: 1
	UserID uint64 `json:"user_id,omitempty"`

	// full_name for user. required for Create
	// required: true
	// min length: 1
	FullName string `json:"full_name,omitempty"`

	ThumnailURL  string            `json:"thumnail_url,omitempty"`
	Technologies []Technology      `json:"technologies,omitempty"`
	Fields       []Field           `json:"fields,omitempty"`
	Contact      map[string]string `json:"contact,omitempty"`
	Description  string            `json:"description,omitempty"`
}

type FilterSortPage struct {
	Filter map[string]string
	Sort   map[string]string
	Page   uint64
}

type ProfileRepository interface {
	InitData() error
	Create(*Profile) error
	Get(uint64) (*Profile, error)
	Update(*Profile) error
	GetFromOffsetToLimitOfProfile(int, int) ([]Profile, error)
	GetWithLimitLastID(int, uint64) ([]Profile, error)
	FilterProfileByFields([]uint64) ([]Profile, error)
	FilterProfileByTechs([]uint64) ([]Profile, error)
	FilterProfileByFieldsTechs([]uint64, []uint64) ([]Profile, error)
	GetMentorWithFilterSortPage(*FilterSortPage) ([]Profile, *uint64, error)
	// Delete(*Profile) error
}

func FromStringToSqlNullString(str string) sql.NullString {
	result := sql.NullString{}
	if str != "" {
		result.Valid = true
		result.String = str
	} else {
		result.Valid = false
	}

	return result
}

func ContactToString(contact map[string]string) (string, error) {
	// from map using json.marsahll -> json ([]byte)-> string
	var strContact string
	b, err := json.Marshal(contact)
	if err != nil {
		return "", err
	}
	strContact = string(b)
	return strContact, nil
}

func FromStringToContact(strContact string) (map[string]string, error) {

	b := []byte(strContact)
	var contact map[string]string
	err := json.Unmarshal(b, &contact)
	if err != nil {
		return nil, err
	}
	return contact, nil

}

func (pJSON *ProfileJSON) ToProfile() (*Profile, error) {
	descritption := FromStringToSqlNullString(pJSON.Description)
	contact, err := ContactToString(pJSON.Contact)
	if err != nil {
		return nil, err
	}

	return &Profile{
		UserID:       pJSON.UserID,
		FullName:     pJSON.FullName,
		ThumnailURL:  pJSON.ThumnailURL,
		Technologies: pJSON.Technologies,
		Fields:       pJSON.Fields,
		Contact:      contact,      //from map[string]string to string
		Description:  descritption, //from string to sql.NullString
	}, nil
}

func (profile *Profile) ToProfileJSON() (*ProfileJSON, error) {
	contact, err := FromStringToContact(profile.Contact)
	if err != nil {
		return nil, err
	}
	return &ProfileJSON{
		UserID:       profile.UserID,
		FullName:     profile.FullName,
		ThumnailURL:  profile.ThumnailURL,
		Technologies: profile.Technologies,
		Fields:       profile.Fields,
		Contact:      contact,
		Description:  profile.Description.String,
	}, nil

}
