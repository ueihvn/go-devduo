package service

import (
	"github.com/ueihvn/go-devduo/model"
	"gorm.io/gorm"
)

type ProfileDb struct {
	Db *gorm.DB
}

var _ model.ProfileRepository = &ProfileDb{}

func NewProfileDb(db *gorm.DB) *ProfileDb {
	return &ProfileDb{
		Db: db,
	}
}

func NewProfileRepository(db *gorm.DB) *ProfileDb {
	return NewProfileDb(db)
}

func (profileDb *ProfileDb) Create(profile interface{}) error {
	err := profileDb.Db.Create(&profile).Error
	if err != nil {
		return err
	}

	return nil
}
