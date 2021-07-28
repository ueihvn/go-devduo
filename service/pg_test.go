package service

import (
	"testing"

	"github.com/ueihvn/go-devduo/model"
)

func TestSetupDb(t *testing.T) {
	db, err := SetupDb()
	if err != nil {
		t.Error(err)
	}

	userDb := NewUserDb(db)

	user := model.User{
		FullName: "full name user test",
		UserName: "User name user test",
		Password: "password user test",
		Email:    "emial user test",
	}
	err = userDb.Create(user)
	if err != nil {
		t.Error(err)
	}

}
