package service

import (
	"errors"
	"strings"

	"github.com/ueihvn/go-devduo/model"
	"gorm.io/gorm"
)

type UserDb struct {
	Db *gorm.DB
}

var _ model.UserRepository = &UserDb{}

func NewUserRepository(db *gorm.DB) *UserDb {
	return NewUserDb(db)
}

func NewUserDb(db *gorm.DB) *UserDb {
	return &UserDb{
		Db: db,
	}
}

func (userDb *UserDb) CreateUser(user *model.User) error {
	err := userDb.Db.Create(&user).Error
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return errors.New("duplicate email or username")
		}
		return errors.New("database error")
	}

	return nil
}

func (userDb *UserDb) GetUserById(userId uint64) (*model.User, error) {
	var user model.User
	err := userDb.Db.First(&user, userId).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userDb *UserDb) GetUserByEmail(email string) (*model.User, error) {
	var user model.User
	err := userDb.Db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (userDb *UserDb) GetAllUsers() ([]model.User, error) {
	var users []model.User
	err := userDb.Db.Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (userDb *UserDb) UpdateUser(user *model.User) error {
	err := userDb.Db.Model(&user).Updates(user).Error
	if err != nil {
		if strings.Contains(err.Error(), "WHERE") {
			return errors.New("missing user ID")
		}
		return errors.New("database err")
	}
	return nil
}

func (userDb *UserDb) InitData() error {

	users := []model.User{
		{
			Password: "passworduser1",
			Email:    "email1@gmail.com",
		},
		{
			Password: "passworduser2",
			Email:    "email2@gmail.com",
		},
		{
			Password: "passworduser3",
			Email:    "email3@gmail.com",
		},
		{
			Password: "passworduser4",
			Email:    "email3@gmail.com",
		},
		{
			Password: "passworduser5",
			Email:    "email5@gmail.com",
		},
	}

	for _, user := range users {
		err := userDb.CreateUser(&user)
		if err != nil {
			return err
		}

	}
	return nil

}
