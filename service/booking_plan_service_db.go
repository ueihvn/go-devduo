package service

import (
	"github.com/ueihvn/go-devduo/model"
	"gorm.io/gorm"
)

type BookingPlanServiceDb struct {
	Db *gorm.DB
}

func NewBookingPlanServiceDb(db *gorm.DB) *BookingPlanServiceDb {
	return &BookingPlanServiceDb{
		Db: db,
	}
}

func NewBookingPlanServiceRepository(db *gorm.DB) *BookingPlanServiceDb {
	return NewBookingPlanServiceDb(db)
}

func (bps *BookingPlanServiceDb) Create(bookingPlanService *model.BookingPlanService) error {
	err := bps.Db.Create(bookingPlanService).Error
	if err != nil {
		return err
	}
	return nil
}

func (bps *BookingPlanServiceDb) Get(id uint64) (*model.BookingPlanService, error) {
	var bookingPlanService model.BookingPlanService
	err := bps.Db.First(&bookingPlanService, id).Error
	if err != nil {
		return nil, err
	}

	return &bookingPlanService, nil
}

func (bps *BookingPlanServiceDb) Update(bookingPlanService *model.BookingPlanService) error {
	err := bps.Db.Model(bookingPlanService).Updates(bookingPlanService).Error
	if err != nil {
		return err
	}
	return nil
}

func (bps *BookingPlanServiceDb) CountUserBookPlanServiceByUserID(userID uint64) (uint64, error) {
	var count int64

	err := bps.Db.Model(&model.BookingPlanService{}).
		Joins("inner join plan_services on booking_plan_services.plan_service_id = plan_services.id").
		Distinct("booking_plan_services.user_id").
		Where("plan_services.user_id = ?", userID).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return uint64(count), nil
}
