package user

import (
	"github.com/AlexRipoll/enchante_technical_interview/pkg/crypto"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldCreateValidAccount(t *testing.T) {
	account, err := New(
		"34d6d340-271a-4db2-b29f-21211148854b",
		"john&Doe ",
		" johndoe@gmail.com ",
		"pass123")

	assert.Nil(t, err, "It shouldn't be any error")
	assert.NotNil(t, account, "account shouldn't be null")
	assert.IsType(t, &Account{}, account, "It should return a account instance")
	assert.Equal(t, "john&amp;Doe", account.Username)
	assert.Equal(t, "johndoe@gmail.com", account.Email)
	assert.Nil(t, crypto.Bcrypt().CheckHash("pass123", account.Password))
	assert.Equal(t, "user", account.Role)
}

func TestShouldReturnErrorWhenInvalidId(t *testing.T) {
	account, err := New(
		"34d6d340-271a-4db2-b29f-21211148854",
		"john&Doe ",
		" johndoe@gmail.com ",
		"pass123")

	assert.Nil(t, account)
	assert.NotNil(t, err)
	assert.IsType(t, &errors.Rest{}, err)
	assert.Equal(t, "invalid uuid", err.Message)
}

func TestShouldReturnErrorWhenNullUsername(t *testing.T) {
	account, err := New(
		"34d6d340-271a-4db2-b29f-21211148854b",
		"",
		" johndoe@gmail.com ",
		"pass123")

	assert.Nil(t, account)
	assert.NotNil(t, err)
	assert.IsType(t, &errors.Rest{}, err)
	assert.Equal(t, "username can't be null", err.Message)
}

func TestShouldReturnErrorWhenUsernameTooShort(t *testing.T) {
	account, err := New(
		"34d6d340-271a-4db2-b29f-21211148854b",
		"jd",
		" johndoe@gmail.com",
		"pass123")

	assert.Nil(t, account)
	assert.NotNil(t, err)
	assert.IsType(t, &errors.Rest{}, err)
	assert.Equal(t, "username must be at least 4 characters long", err.Message)
}

func TestShouldReturnErrorWhenNullEmail(t *testing.T) {
	account, err := New(
		"34d6d340-271a-4db2-b29f-21211148854b",
		"johnDoe",
		"",
		"pass123")

	assert.Nil(t, account)
	assert.NotNil(t, err)
	assert.IsType(t, &errors.Rest{}, err)
	assert.Equal(t, "invalid email address", err.Message)
}

func TestShouldReturnErrorWhenInvalidEmail(t *testing.T) {
	account, err := New(
		"34d6d340-271a-4db2-b29f-21211148854b",
		"johnDoe",
		"john&doe@gmail.co",
		"pass123")

	assert.Nil(t, account)
	assert.NotNil(t, err)
	assert.IsType(t, &errors.Rest{}, err)
	assert.Equal(t, "invalid email address", err.Message)
}

func TestShouldReturnErrorWhenNullPassword(t *testing.T) {
	account, err := New(
		"34d6d340-271a-4db2-b29f-21211148854b",
		"johnDoe",
		"johndoe@gmail.com",
		"")

	assert.Nil(t, account)
	assert.NotNil(t, err)
	assert.IsType(t, &errors.Rest{}, err)
	assert.Equal(t, "password can't be null", err.Message)
}

func TestShouldReturnErrorWhenPasswordTooShort(t *testing.T) {
	account, err := New(
		"34d6d340-271a-4db2-b29f-21211148854b",
		"johnDoe",
		"johndoe@gmail.com",
		"pass1")

	assert.Nil(t, account)
	assert.NotNil(t, err)
	assert.IsType(t, &errors.Rest{}, err)
	assert.Equal(t, "password must contain at least 6 characters", err.Message)
}
