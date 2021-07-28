package service

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ueihvn/go-devduo/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPg(DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
	dburl := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		DbHost,
		DbUser,
		DbPassword,
		DbName,
		DbPort,
	)

	db, err := gorm.Open(postgres.Open(dburl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SetupDb() (*gorm.DB, error) {
	err := godotenv.Load("../.env.development")
	if err != nil {
		return nil, err
	}

	DbUser := os.Getenv("POSTGRES_USER")
	DbPassword := os.Getenv("POSTGRES_PASSWORD")
	DbPort := os.Getenv("POSTGRES_PORT")
	DbHost := os.Getenv("POSTGRES_HOST")
	DbName := os.Getenv("POSTGRES_DB")

	db, err := NewPg(DbUser, DbPassword, DbPort, DbHost, DbName)
	if err != nil {
		return nil, err
	}

	db.Migrator().DropTable(
		&model.ProgrammingLanguage{},
		&model.Field{},
		&model.User{},
		&model.Profile{},
		&model.PlanService{},
		&model.BookingPlanService{},
	)
	err = db.AutoMigrate(
		&model.ProgrammingLanguage{},
		&model.Field{},
		&model.User{},
		&model.Profile{},
		&model.PlanService{},
	)

	if err != nil {
		return nil, errors.New("err automigrate")
	}

	err = db.SetupJoinTable(&model.User{}, "BookingPlanServices", &model.BookingPlanService{})

	if err != nil {
		return nil, errors.New("error setup join table")
	}

	return db, nil
}
