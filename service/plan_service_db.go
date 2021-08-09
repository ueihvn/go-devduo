package service

import (
	"errors"
	"strings"

	"github.com/ueihvn/go-devduo/model"
	"gorm.io/gorm"
)

type PlanServiceDb struct {
	Db *gorm.DB
}

func NewPlanServiceDb(db *gorm.DB) *PlanServiceDb {
	return &PlanServiceDb{
		Db: db,
	}
}

func NewPlanServiceRepository(db *gorm.DB) *PlanServiceDb {
	return NewPlanServiceDb(db)
}

func (ps *PlanServiceDb) Create(planService *model.PlanService) error {
	err := ps.Db.Create(planService).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return errors.New("duplicate title service")
		}
		return err
	}
	return nil
}

func (ps *PlanServiceDb) Get(id uint64) (*model.PlanService, error) {
	var planService model.PlanService
	err := ps.Db.First(&planService, id).Error
	if err != nil {
		return nil, err
	}

	return &planService, nil
}

func (ps *PlanServiceDb) Update(planService *model.PlanService) error {
	err := ps.Db.Model(planService).Updates(planService).Error
	if err != nil {
		return err
	}
	return nil
}
