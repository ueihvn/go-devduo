package service

import (
	"fmt"
	"testing"

	"github.com/ueihvn/go-devduo/model"
)

// test pg.go
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

func TestInitData(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error(err)
	}

	err = InitData(db)
	if err != nil {
		t.Error(err)
	}
}

// test user_db.go
func TestCreateUser(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}

	userDb := NewUserRepository(db)

	user := model.User{
		FullName: "full name user test",
		UserName: "Username1",
		Password: "passwordusertest",
		Email:    "email1@gmail.com",
	}
	err = userDb.CreateUser(&user)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(user)
	}
}

func TestUpdateUser(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}

	userDb := NewUserRepository(db)

	user := model.User{
		ID:       1,
		Password: "doipasswordroine",
	}
	err = userDb.UpdateUser(&user)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(user)
	}
}

//test profile_db.go
func TestCreateProfile(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}

	profileDb := NewProfileDb(db)

	pJSON := model.ProfileJSON{
		UserID: 2,
		Technologies: []model.Technology{
			{
				ID:   1,
				Name: "C",
			},
			{
				ID:   2,
				Name: "C++",
			},
		},
		Fields: []model.Field{
			{
				ID:   3,
				Name: "Blockchain",
			},
			{
				ID:   4,
				Name: "IoT",
			},
		},
		Contact:     map[string]string{},
		Description: "test description ne ong oi",
	}
	profile, err := pJSON.ToProfile()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", profile)

	err = profileDb.Create(profile)
	if err != nil {
		t.Error(err)
	}

}

func TestGetProfile(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}
	profileDb := NewProfileDb(db)

	// getfields
	fieldDb := NewFieldRepository(db)
	fields, err := fieldDb.GetFieldsByUserId(1)
	if err != nil {
		t.Errorf("err profileDb.GetFieldsByUserId - err: %s", err)
	}

	// gettechs
	techDb := NewTechnologyRepository(db)
	techs, err := techDb.GetTechnologiesByUserId(1)
	if err != nil {
		t.Errorf("err technologyDb.GetTechnologiesByUserId - err: %s", err)
	}

	profile, err := profileDb.Get(1)
	if err != nil {
		t.Errorf("error profileDb.Get- err: %s", err)
	}

	profile.Fields = fields
	profile.Technologies = techs
	fmt.Printf("%+v\n", profile)
}

func TestUpdateProfile(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}
	profileDb := NewProfileDb(db)

	pJSON := model.ProfileJSON{
		UserID: 1,
		Technologies: []model.Technology{
			{
				ID:   7,
				Name: "Go",
			},
			{
				ID:   8,
				Name: "Python",
			},
		},
		Fields: []model.Field{
			{
				ID:   3,
				Name: "Blockchain",
			},
			{
				ID:   1,
				Name: "AI",
			},
		},
		Contact:     map[string]string{},
		Description: "test description ne ong oi",
	}

	profile, err := pJSON.ToProfile()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", profile)
	err = profileDb.Update(profile)
	if err != nil {
		t.Error(err)
	}

}

// test field_fb.go
func TestGetFieldsByUserId(t *testing.T) {

	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}

	fieldDb := NewFieldRepository(db)
	fields, err := fieldDb.GetFieldsByUserId(1)
	if err != nil {
		t.Errorf("err profileDb.GetFieldsByUserId - err: %s", err)
	}
	fmt.Println(fields)
}

// test technology_db.go
func TestGetTechnologiesByUserId(t *testing.T) {

	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}

	techDb := NewTechnologyRepository(db)
	techs, err := techDb.GetTechnologiesByUserId(1)
	if err != nil {
		t.Errorf("err technologyDb.GetTechnologiesByUserId - err: %s", err)
	}
	fmt.Println(techs)
}
