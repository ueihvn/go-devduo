package service

import (
	"errors"
	"fmt"

	"github.com/ueihvn/go-devduo/config"
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
	dbConfig := config.NewConfig()
	db, err := NewDb(
		dbConfig.PostgresUser,
		dbConfig.PostgresPassword,
		dbConfig.PostgrePort,
		dbConfig.PostgresHost,
		dbConfig.PostgresDb,
	)
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
		&model.BookingPlanService{},
	)

	if err != nil {
		return errors.New("err automigrate")
	}

	return nil
}

func InitData(db *gorm.DB) error {

	tDb := NewTechnologyRepository(db)
	err := tDb.InitData()
	if err != nil {
		return errors.New("err tech InitData")
	}

	fDb := NewFieldRepository(db)
	err = fDb.InitData()
	if err != nil {
		return errors.New("err field InitData")
	}

	return nil
}
