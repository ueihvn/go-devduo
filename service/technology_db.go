package service

import (
	"errors"
	"strings"

	"github.com/ueihvn/go-devduo/model"
	"gorm.io/gorm"
)

type TechnologyDb struct {
	Db *gorm.DB
}

func NewTechnologyRepository(db *gorm.DB) *TechnologyDb {
	return &TechnologyDb{
		Db: db,
	}
}

func (technologyDb *TechnologyDb) Create(technology *model.Technology) error {
	err := technologyDb.Db.Create(technology).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return errors.New("duplicate technology")
		}
		return errors.New("database error")
	}
	return nil
}

func (technologyDb *TechnologyDb) GetAll() ([]model.Technology, error) {
	var technologies []model.Technology
	err := technologyDb.Db.Find(&technologies).Error
	if err != nil {
		return nil, err
	}

	return technologies, nil
}

func (technologyDb *TechnologyDb) Get(techName string) (*model.Technology, error) {
	var technology model.Technology
	err := technologyDb.Db.Where("name = ?", techName).Find(&technology).Error
	if err != nil {
		return nil, err
	}

	return &technology, nil
}

func (technologyDb *TechnologyDb) GetTechnologiesByUserId(userId int) ([]model.Technology, error) {
	var techs []model.Technology
	err := technologyDb.Db.Table("technologies").
		Select("technologies.name,technologies.id").
		Joins("left join profile_technologies on technologies.id = profile_technologies.technology_id").
		Where("profile_technologies.profile_user_id = ?", userId).Find(&techs).Error
	if err != nil {
		return nil, err
	}

	return techs, nil
}

func (technologyDb *TechnologyDb) Delete(technology *model.Technology) error {
	err := technologyDb.Db.Delete(technology).Error
	if err != nil {
		return err
	}

	return nil
}

func (technologyDb *TechnologyDb) InitData() error {
	strTech := []string{
		"C",
		"C++",
		"C#",
		"Java",
		"PHP",
		"Ruby",
		"Go",
		"Python",
		"Javascript",
		"HTML",
		"CSS",
		"ReactJS",
		"React Native",
		"Spring boot",
		".NET",
		"Arduino",
		"Django",
		"NodeJS",
		"Angular",
		"VueJS",
	}

	technologies := make([]model.Technology, len(strTech))
	for index, techName := range strTech {
		tech := model.Technology{
			Name: techName,
		}
		technologies[index] = tech
	}
	err := technologyDb.Db.Create(&technologies).Error
	if err != nil {
		return err
	}

	return nil
}
