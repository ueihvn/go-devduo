package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
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
func (ps *PlanServiceDb) GetPlanServiceByUserID(userId uint64) ([]model.PlanService, error) {
	var planServices []model.PlanService
	err := ps.Db.Where("user_id = ?", userId).Find(&planServices).Error
	if err != nil {
		return nil, err
	}

	return planServices, err
}

func (ps *PlanServiceDb) GetSmallestPricePlanServiceByUserID(userID uint64) (*decimal.Decimal, error) {
	result := struct {
		Min decimal.Decimal
	}{
		Min: decimal.Decimal{},
	}
	err := ps.Db.Model(&model.PlanService{}).Select("user_id, MIN(price)").Group("user_id").Having("user_id = ?", userID).Find(&result).Error
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v\n", result)

	// res, _ := decimal.NewFromString(result.min)
	return &result.Min, nil
}
