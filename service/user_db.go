package service

import (
	"github.com/ueihvn/go-devduo/model"
	"gorm.io/gorm"
)

type UserDb struct {
	Db *gorm.DB
}

func NewUserDb(db *gorm.DB) *UserDb {
	return &UserDb{
		Db: db,
	}
}

func (userDb *UserDb) Create(user model.User) error {
	result := userDb.Db.Create(&user)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (userDb *UserDb) RetrieveById(userId uint) (*model.User, error) {
	var user model.User
	result := userDb.Db.First(&user, userId)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}
