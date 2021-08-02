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

func NewDb(DbUser, DbPassword, DbPort, DbHost, DbName string) (*gorm.DB, error) {
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

func ConnectDb() (*gorm.DB, error) {
	err := godotenv.Load(".env.development")
	if err != nil {
		return nil, err
	}

	DbUser := os.Getenv("POSTGRES_USER")
	DbPassword := os.Getenv("POSTGRES_PASSWORD")
	DbPort := os.Getenv("POSTGRES_PORT")
	DbHost := os.Getenv("POSTGRES_HOST")
	DbName := os.Getenv("POSTGRES_DB")

	db, err := NewDb(DbUser, DbPassword, DbPort, DbHost, DbName)
	if err != nil {
		return nil, errors.New("fail connect to Db")
	}

	return db, nil
}

func MigrateDb(db *gorm.DB) error {

	db.Migrator().DropTable(
		&model.Technology{},
		&model.Field{},
		&model.User{},
		&model.Profile{},
		&model.PlanService{},
		&model.BookingPlanService{},
	)
	err := db.AutoMigrate(
		&model.Technology{},
		&model.Field{},
		&model.User{},
		&model.Profile{},
		&model.PlanService{},
	)

	if err != nil {
		return errors.New("err automigrate")
	}

	err = db.SetupJoinTable(&model.User{}, "BookingPlanServices", &model.BookingPlanService{})

	if err != nil {
		return errors.New("error setup join table")
	}
	return nil
}
