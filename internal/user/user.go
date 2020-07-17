package user

import (
	"fmt"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/crypto"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/uuidv4"
	"html"
	"regexp"
	"strings"
)

const (
	minUsernameLength = 4
	minPasswordLength = 6
	defaultRole = "user"
	emailPattern = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

type Account struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password,omitempty"`
	Role      string `json:"role,omitempty"`
	CreatedOn string `json:"created_on,omitempty"`
}

func New(id string, username string, email string, password string) (*Account, *errors.Rest) {
	a := Account{
		Id:       id,
		Username: username,
		Email:    email,
		Password: password,
	}
	a.Role = defaultRole

	if err := a.validateUsername(); err != nil {
		return nil, err
	}
	if err := a.validateEmail(); err != nil {
		return nil, err
	}
	if err := a.validatePassword(); err != nil {
		return nil, err
	}
	var hashErr error
	a.Password, hashErr = crypto.Bcrypt().Hash(a.Password)
	if hashErr != nil {
		return nil, errors.NewInternalServerError(hashErr.Error())
	}

	return &a, nil
}

func (a *Account) validateUsername() *errors.Rest {
	if err := uuidv4.NewService().Validate(a.Id); err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	a.Username = html.EscapeString(strings.TrimSpace(a.Username))
	if a.Username == "" {
		return errors.NewBadRequestError("username can't be null")
	}
	if len(a.Username) < minUsernameLength {
		return errors.NewBadRequestError(
			fmt.Sprintf("username must be at least %d characters long", minUsernameLength))
	}
	return nil
}

func (a *Account) validateEmail() *errors.Rest {
	a.Email = html.EscapeString(strings.TrimSpace(a.Email))
	regularExpression := regexp.MustCompile(emailPattern)

	if ok := regularExpression.MatchString(a.Email); !ok {
		return errors.NewBadRequestError("invalid email address")
	}
	return nil
}

func (a *Account) validatePassword() *errors.Rest {
	if a.Password == "" {
		return errors.NewBadRequestError("password can't be null")
	}
	if len(a.Password) < minPasswordLength {
		return errors.NewBadRequestError(fmt.Sprintf(
			"password must contain at least %d characters", minPasswordLength))
	}
	return nil
}
