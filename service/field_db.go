package service

import (
	"errors"
	"strings"

	"github.com/ueihvn/go-devduo/model"
	"gorm.io/gorm"
)

type FieldDb struct {
	Db *gorm.DB
}

func NewFieldRepository(db *gorm.DB) *FieldDb {
	return &FieldDb{
		Db: db,
	}
}

func (fieldDb *FieldDb) Create(field *model.Field) error {
	err := fieldDb.Db.Create(field).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return errors.New("duplicate field")
		}
		return errors.New("database error")
	}
	return nil
}

func (fieldDb *FieldDb) GetAll() ([]model.Field, error) {
	var fields []model.Field
	err := fieldDb.Db.Find(&fields).Error
	if err != nil {
		return nil, err
	}

	return fields, nil
}

func (fieldDb *FieldDb) Get(fieldName string) (*model.Field, error) {
	var field model.Field
	err := fieldDb.Db.Where("name = ?", fieldName).Find(&field).Error
	if err != nil {
		return nil, err
	}

	return &field, nil
}

func (fieldDb *FieldDb) GetFieldsByUserId(userId uint64) ([]model.Field, error) {
	var fields []model.Field
	err := fieldDb.Db.Table("fields").
		Select("fields.name,fields.id").
		Joins("left join profile_fields on fields.id = profile_fields.field_id").
		Where("profile_fields.profile_user_id = ?", userId).Find(&fields).Error
	if err != nil {
		return nil, err
	}

	return fields, nil
}

func (fieldDb *FieldDb) Delete(field *model.Field) error {
	err := fieldDb.Db.Delete(field).Error
	if err != nil {
		return err
	}

	return nil
}

func (fieldDb *FieldDb) InitData() error {
	strFields := []string{
		"AI",
		"Big Data",
		"Blockchain",
		"IoT",
		"ML",
		"DL",
		"Network",
		"E-commerce",
		"DevOps",
		"System",
		"Mobile App",
		"Web App",
		"Desktop App",
	}

	fields := make([]model.Field, len(strFields))
	for index, fieldName := range strFields {
		field := model.Field{
			Name: fieldName,
		}
		fields[index] = field
	}
	err := fieldDb.Db.Create(&fields).Error
	if err != nil {
		return err
	}

	return nil
}
