package service

import (
	"fmt"
	"testing"

	"github.com/ueihvn/go-devduo/model"
)

func TestConnectDb(t *testing.T) {
	_, err := ConnectDb()
	if err != nil {
		t.Error(err)
	}
}

func TestMigrateDb(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error(err)
	}

	err = MigrateDb(db)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateUser(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}

	userDb := NewUserDb(db)

	user := model.User{
		FullName: "full name user test",
		UserName: "Usernameusertest",
		Password: "passwordusertest",
		Email:    "email3@gmail.com",
	}
	err = userDb.CreateUser(&user)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(user)
	}
}
