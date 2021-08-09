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

func (profileDb *ProfileDb) Create(profile *model.Profile) error {
	err := profileDb.Db.Omit("Technologies.*", "Fields.*").
		Preload("Technologies").Preload("Fields").
		Create(&profile).Error
	if err != nil {
		return err
	}

	return nil
}

func (profileDb *ProfileDb) Get(userId uint64) (*model.Profile, error) {
	var profile model.Profile
	err := profileDb.Db.Where("user_id", userId).First(&profile).Error
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (profileDb *ProfileDb) Update(profile *model.Profile) error {
	err := profileDb.Db.Model(&profile).Updates(profile).Error
	if err != nil {
		return err
	}
	// update techs, fields
	err = profileDb.Db.Model(&profile).Association("Technologies").Replace(profile.Technologies)

	if err != nil {
		return err
	}

	err = profileDb.Db.Model(&profile).Association("Fields").Replace(profile.Fields)

	if err != nil {
		return err
	}

	return nil
}
