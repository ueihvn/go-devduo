package service

import (
	"errors"
	"strings"

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
		if strings.Contains(err.Error(), "not found") {
			return nil, errors.New("user no found")
		}
		return nil, err
	}

	var techs []model.Technology
	err = profileDb.Db.Table("technologies").
		Select("technologies.name,technologies.id").
		Joins("left join profile_technologies on technologies.id = profile_technologies.technology_id").
		Where("profile_technologies.profile_user_id = ?", userId).Find(&techs).Error
	if err != nil {
		return nil, err
	}

	var fields []model.Field
	err = profileDb.Db.Table("fields").
		Select("fields.name,fields.id").
		Joins("left join profile_fields on fields.id = profile_fields.field_id").
		Where("profile_fields.profile_user_id = ?", userId).Find(&fields).Error
	if err != nil {
		return nil, err
	}
	profile.Technologies = techs
	profile.Fields = fields

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

func (profileDb *ProfileDb) GetFromOffsetToLimitOfProfile(offset, limit int) ([]model.Profile, error) {
	var profiles []model.Profile

	if offset == 0 {
		err := profileDb.Db.Limit(limit).Find(&profiles).Error
		if err != nil {
			return nil, err
		}
		return profiles, nil
	}

	err := profileDb.Db.Limit(limit).Offset(offset).Find(&profiles).Error
	if err != nil {
		return nil, err
	}

	return profiles, nil

}

func (profileDb *ProfileDb) GetWithLimitLastID(limit int, last_id uint64) ([]model.Profile, error) {

	var profiles []model.Profile

	err := profileDb.Db.Limit(limit).Where("user_id > ?", last_id).Find(&profiles).Error
	if err != nil {
		return nil, err
	}

	return profiles, nil

}
