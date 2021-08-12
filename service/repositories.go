package service

import "github.com/ueihvn/go-devduo/model"

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

	return &Repositories{
		Ur:   NewUserRepository(db),
		Pr:   NewProfileRepository(db),
		Fr:   NewFieldRepository(db),
		Tr:   NewTechnologyRepository(db),
		Psr:  NewPlanServiceRepository(db),
		Bpsr: NewBookingPlanServiceRepository(db),
	}, nil
}
