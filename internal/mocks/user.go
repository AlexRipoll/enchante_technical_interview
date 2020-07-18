package mocks

import (
	"fmt"
	"github.com/AlexRipoll/enchante_technical_interview/internal/user"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
)

var (
	userRepositoryData = []user.Account{
		{
			Id:       "34d6d340-271a-4db2-b29f-21211148854b",
			Username: "johnDoe",
			Email:    "johndoe@gmail.com",
			Password: "$2a$10$RqYsU9IFiCqu.zG1f3yMoeYdPESQlSYm1AhlGVoSbrjkD4mpBH.iS"},
		{
			Id:       "652ef79a-8cc4-4660-a0d2-21e7547249bc",
			Username: "janeDoe",
			Email:    "janedoe@gmail.com",
			Password: "$2a$10$W7vAj3SK2J6S9eG5RIoLyufVYd5cuEnG32d7UAwkjyASmDOh8HRRW"},
	}
)

type userRepository struct {
}

func UserRepository() user.Repository {
	return &userRepository{}
}

func (r *userRepository) Find(id string) (*user.Account, *errors.Rest) {
	if id == userRepositoryData[0].Id {
		return &userRepositoryData[0], nil
	}
	return nil, errors.NewNotFoundError(fmt.Sprintf("no user found with id %s", id))
}

func (r *userRepository) Save(account *user.Account) *errors.Rest {
	return nil
}

func (r *userRepository) Delete(id string) *errors.Rest {
	return nil
}

func (r *userRepository) FindByEmail(email string) (*user.Account, *errors.Rest) {
	if email != userRepositoryData[0].Email {
		return nil, nil
	}
	return &userRepositoryData[0], nil
}
