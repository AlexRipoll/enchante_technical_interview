package user_test

import (
	"github.com/AlexRipoll/enchante_technical_interview/internal/mocks"
	"github.com/AlexRipoll/enchante_technical_interview/internal/user"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldRegisterUser(t *testing.T) {
	err := user.NewService(mocks.UserRepository()).Register(
		"jane",
		"janedoe@gmail.com",
		"pass123")

	assert.Nil(t, err)
}

func TestShouldNotRegisterUserAccountWhenBadFieldsProvided(t *testing.T) {
	err := user.NewService(mocks.UserRepository()).Register(
		"jane",
		"jane&doe@gmail.com",
		"pass123")

	assert.NotNil(t, err)
	assert.IsType(t, &errors.Rest{}, err)
	assert.Equal(t, "invalid email address", err.Message)
}

func TestShouldNotRegisterUserAccountIfEmailAlreadyTaken(t *testing.T) {
	err := user.NewService(mocks.UserRepository()).Register(
		"johnDoe",
		"johndoe@gmail.com",
		"pass123")

	assert.NotNil(t, err)
	assert.IsType(t, &errors.Rest{}, err)
	assert.Equal(t, "email is already registered", err.Message)
}

func TestShouldFindUserAccount(t *testing.T) {
	u, err := user.NewService(mocks.UserRepository()).Search("34d6d340-271a-4db2-b29f-21211148854b")

	assert.Nil(t, err)
	assert.NotNil(t, u)
	assert.IsType(t, &user.Account{}, u)
}

func TestShouldNotFindUserAccount(t *testing.T) {
	u, err := user.NewService(mocks.UserRepository()).Search("34d6d340-271a-4db2-b29f-212111488547")

	assert.Nil(t, u)
	assert.NotNil(t, err)
	assert.IsType(t, &errors.Rest{}, err)
	assert.Equal(t, "no user found with id 34d6d340-271a-4db2-b29f-212111488547", err.Message)
}

func TestShouldDeleteUserAccount(t *testing.T) {
	err := user.NewService(mocks.UserRepository()).Delete("34d6d340-271a-4db2-b29f-21211148854b")

	assert.Nil(t, err)
}

func TestShouldNotDeleteUserAccount(t *testing.T) {
	err := user.NewService(mocks.UserRepository()).Delete("34d6d340-271a-4db2-b29f-212111488547")

	assert.NotNil(t, err)
	assert.IsType(t, &errors.Rest{}, err)
	assert.Equal(t, "no user found with id 34d6d340-271a-4db2-b29f-212111488547", err.Message)
}
