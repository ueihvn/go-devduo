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

func TestGetUserByEmail(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}

	userDb := NewUserRepository(db)
	email := "email4@gmail.com"
	user, err := userDb.GetUserByEmail(email)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", user)
}

func TestGetUserById(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}

	userDb := NewUserRepository(db)
	user, err := userDb.GetUserById(6)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v\n", user)
}

//test profile_db.go
func TestCreateProfile(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}

	profileDb := NewProfileDb(db)

	pJSON := model.ProfileJSON{
		UserID:   2,
		FullName: "full name ne",
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
	userID := uint64(1)

	profile, err := profileDb.Get(userID)
	if err != nil {
		t.Errorf("error profileDb.Get- err: %s", err)
	}

	fmt.Printf("%+v\n", profile)
}

func TestGetTest(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}
	profileDb := NewProfileDb(db)

	profile, err := profileDb.GetTest()
	if err != nil {
		t.Errorf("error profileDb.Get- err: %s", err)
	}

	fmt.Printf("%+v\n", profile)
}

func TestUpdateProfile(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}
	profileDb := NewProfileDb(db)

	pJSON := model.ProfileJSON{
		UserID:   1,
		FullName: "full name nua ne",
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

func TestGetProfileOL(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}
	profileDb := NewProfileDb(db)

	profiles, err := profileDb.GetFromOffsetToLimitOfProfile(0, 50)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(profiles)

}

func TestFilterProfileByField(t *testing.T) {

	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}
	profileDb := NewProfileDb(db)

	fields := []uint64{3, 4}
	profiles, err := profileDb.FilterProfileByFields(fields)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(profiles)

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

// test plan_service.go
func TestCreatePS(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}

	psDb := NewPlanServiceRepository(db)
	sp := model.PlanService{
		Title: "debug ko 1",
		Price: model.PriceStringToDecimal("100000"),
	}

	err = psDb.Create(&sp)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sp)
}

func TestGetSmallestPricePS(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}

	psDb := NewPlanServiceRepository(db)
	sp, err := psDb.GetSmallestPricePlanServiceByUserID(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(sp)
}

func TestGetCountUser(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}

	bpsDb := NewBookingPlanServiceRepository(db)
	count, err := bpsDb.CountUserBookPlanServiceByUserID(1)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(count)

}

// test booking_plan_services
func TestCreateBookingPlanService(t *testing.T) {
	db, err := ConnectDb()
	if err != nil {
		t.Error("error connect db")
	}

	bps := model.BookingPlanService{
		UserID:        3,
		PlanServiceID: 2,
	}
	bpsDb := NewBookingPlanServiceRepository(db)
	err = bpsDb.Create(&bps)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%+v\n", bps)

}
