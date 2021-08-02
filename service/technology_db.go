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

func (technologyDb *TechnologyDb) Delete(technology *model.Technology) error {
	err := technologyDb.Db.Delete(technology).Error
	if err != nil {
		return err
	}

	return nil
}
