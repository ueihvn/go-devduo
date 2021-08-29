package service

import (
	"errors"
	"fmt"
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

	err = profileDb.fillTechsFieldsProfile(&profile)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}

func (profileDb *ProfileDb) fillTechsFieldsProfile(profile *model.Profile) error {
	err := profileDb.Db.Table("technologies").
		Select("technologies.name,technologies.id").
		Joins("left join profile_technologies on technologies.id = profile_technologies.technology_id").
		Where("profile_technologies.profile_user_id = ?", profile.UserID).Find(&profile.Technologies).Error
	if err != nil {
		return err
	}

	err = profileDb.Db.Table("fields").
		Select("fields.name,fields.id").
		Joins("left join profile_fields on fields.id = profile_fields.field_id").
		Where("profile_fields.profile_user_id = ?", profile.UserID).Find(&profile.Fields).Error
	if err != nil {
		return err
	}
	return nil
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
	} else {
		err := profileDb.Db.Limit(limit).Offset(offset).Find(&profiles).Error
		if err != nil {
			return nil, err
		}
	}

	for i := range profiles {
		err := profileDb.fillTechsFieldsProfile(&profiles[i])
		if err != nil {
			return nil, err
		}

	}
	return profiles, nil

}

func (profileDb *ProfileDb) GetWithLimitLastID(limit int, last_id uint64) ([]model.Profile, error) {

	var profiles []model.Profile

	err := profileDb.Db.Limit(limit).Where("user_id > ?", last_id).Find(&profiles).Error
	if err != nil {
		return nil, err
	}

	for i := range profiles {
		err := profileDb.fillTechsFieldsProfile(&profiles[i])
		if err != nil {
			return nil, err
		}

	}
	return profiles, nil

}

func (profileDb *ProfileDb) FilterProfileByFields(fields []uint64) ([]model.Profile, error) {
	var profiles []model.Profile

	err := profileDb.Db.Raw(filterProfileByField, fields).Scan(&profiles).Error
	if err != nil {
		return nil, err
	}

	for i := range profiles {
		err := profileDb.fillTechsFieldsProfile(&profiles[i])
		if err != nil {
			return nil, err
		}

	}

	return profiles, nil
}

func (profileDb *ProfileDb) FilterProfileByTechs(techs []uint64) ([]model.Profile, error) {
	var profiles []model.Profile

	err := profileDb.Db.Raw(filterProfileByTech, techs).Scan(&profiles).Error
	if err != nil {
		return nil, err
	}

	for i := range profiles {
		err := profileDb.fillTechsFieldsProfile(&profiles[i])
		if err != nil {
			return nil, err
		}

	}

	return profiles, nil
}

func (profileDb *ProfileDb) FilterProfileByFieldsTechs(fields, techs []uint64) ([]model.Profile, error) {
	var profiles []model.Profile

	err := profileDb.Db.Raw(filterProfileByFieldTech, fields, techs).Scan(&profiles).Error
	if err != nil {
		return nil, err
	}

	for i := range profiles {
		err := profileDb.fillTechsFieldsProfile(&profiles[i])
		if err != nil {
			return nil, err
		}

	}

	return profiles, nil
}

func (profileDb *ProfileDb) GetMentorWithFilterSortPage(fsp *model.FilterSortPage) ([]model.Profile, *uint64, error) {

	query, args := buildProfileFilterSortPageQuery(fsp)
	fmt.Printf("query: %v\nargs: %+v\n", query, args)

	var profiles []model.Profile

	res := profileDb.Db.Raw(query, args...).Scan(&profiles)
	if res.Error != nil {
		return nil, nil, res.Error
	}

	for i := range profiles {
		err := profileDb.fillTechsFieldsProfile(&profiles[i])
		if err != nil {
			return nil, nil, err
		}
	}
	rows := uint64(res.RowsAffected)

	return profiles, &rows, nil
}

func (profileDb *ProfileDb) InitData() error {
	profiles := []model.ProfileJSON{
		{
			UserID:   1,
			FullName: "full name user1",
			Technologies: []model.Technology{
				{
					ID:   1,
					Name: "C",
				},
				{
					ID:   2,
					Name: "C++",
				},
			},
			Fields: []model.Field{
				{
					ID:   3,
					Name: "Blockchain",
				},
				{
					ID:   4,
					Name: "IoT",
				},
			},
			Contact:     map[string]string{},
			Description: "test description ne ong oi",
		},
		{
			UserID:   2,
			FullName: "full name user2",
			Technologies: []model.Technology{
				{
					ID:   1,
					Name: "C",
				},
				{
					ID:   2,
					Name: "C++",
				},
				{
					ID:   3,
					Name: "C#",
				},
				{
					ID:   4,
					Name: "Java",
				},
				{
					ID:   5,
					Name: "PHP",
				},
				{
					ID:   6,
					Name: "Ruby",
				},
			},
			Fields: []model.Field{
				{
					ID:   3,
					Name: "Blockchain",
				},
				{
					ID:   4,
					Name: "IoT",
				},
				{
					ID:   8,
					Name: "E-commerce",
				},
				{
					ID:   12,
					Name: "Web App",
				},
			},
			Contact:     map[string]string{},
			Description: "test description ne ong oi",
		},
		{
			UserID:   3,
			FullName: "full name user3",
			Technologies: []model.Technology{
				{
					ID:   10,
					Name: "HTML",
				},
				{
					ID:   12,
					Name: "ReactJS",
				},
				{
					ID:   1,
					Name: "C",
				},
				{
					ID:   2,
					Name: "C++",
				},
				{
					ID:   3,
					Name: "C#",
				},
				{
					ID:   4,
					Name: "Java",
				},
				{
					ID:   5,
					Name: "PHP",
				},
				{
					ID:   6,
					Name: "Ruby",
				},
			},
			Fields: []model.Field{
				{
					ID:   3,
					Name: "Blockchain",
				},
				{
					ID:   4,
					Name: "IoT",
				},
				{
					ID:   8,
					Name: "E-commerce",
				},
				{
					ID:   12,
					Name: "Web App",
				},
			},
			Contact:     map[string]string{},
			Description: "test description ne ong oi",
		},
	}

	for _, profileJSON := range profiles {
		profile, err := profileJSON.ToProfile()
		if err != nil {
			return err
		}
		err = profileDb.Create(profile)
		if err != nil {
			return err
		}

	}
	return nil
}
