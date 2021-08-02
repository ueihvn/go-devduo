package service

import "github.com/ueihvn/go-devduo/model"

type Repositories struct {
	Ur model.UserRepository
	Pr model.ProfileRepository
}

func NewRepositories() (*Repositories, error) {
	db, err := ConnectDb()
	if err != nil {
		return nil, err
	}

	return &Repositories{
		Ur: NewUserRepository(db),
		Pr: NewProfileRepository(db),
	}, nil
}
