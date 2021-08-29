package service

import (
	"github.com/ueihvn/go-devduo/model"
)

type Repositories struct {
	Ur   model.UserRepository
	Pr   model.ProfileRepository
	Fr   model.FieldRepository
	Tr   model.TechnologyRepository
	Psr  model.PlanServiceRepository
	Bpsr model.BookingPlanServiceRepository
}

func NewRepositories() (*Repositories, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, err
	}

	err = MigrateDb(db)
	if err != nil {
		return nil, err
	}

	repo := &Repositories{
		Ur:   NewUserRepository(db),
		Pr:   NewProfileRepository(db),
		Fr:   NewFieldRepository(db),
		Tr:   NewTechnologyRepository(db),
		Psr:  NewPlanServiceRepository(db),
		Bpsr: NewBookingPlanServiceRepository(db),
	}

	return repo, nil
}

func (repositories *Repositories) InitData() error {
	err := repositories.Fr.InitData()
	if err != nil {
		return err
	}

	err = repositories.Tr.InitData()
	if err != nil {
		return err
	}

	err = repositories.Ur.InitData()
	if err != nil {
		return err
	}

	err = repositories.Pr.InitData()
	if err != nil {
		return err
	}

	err = repositories.Psr.InitData()
	if err != nil {
		return err
	}

	err = repositories.Bpsr.InitData()
	if err != nil {
		return err
	}
	return nil
}
