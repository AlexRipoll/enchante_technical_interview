package user

import (
	"github.com/AlexRipoll/enchante_technical_interview/pkg/crypto"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/jwt"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/time"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/uuidv4"
	"net/http"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Register(username string, email string, password string) *errors.Rest {
	id, err := uuidv4.NewService().Generate()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	account, accountErr := New(id, username, email, password)
	if accountErr != nil {
		return accountErr
	}
	account.CreatedOn = time.Current()

	searchResult, searchErr := s.repository.FindByEmail(email)
	if searchErr != nil && searchErr.Status != http.StatusNotFound {
		return searchErr
	}
	if searchResult != nil {
		return errors.NewBadRequestError("email is already registered")
	}

	if dbErr := s.repository.Save(account); dbErr != nil {
		return dbErr
	}
	return nil
}

func (s *service) Login(email, password string) (string, *errors.Rest) {
	account, err := ValidateCredentials(email, password)
	if err != nil {
		return "", err
	}

	account, err = s.repository.FindByEmail(account.Email)
	if err != nil {
		return "", errors.NewNotFoundError("no user found with the given credentials")
	}
	if err := crypto.Bcrypt().CheckHash(password, account.Password); err != nil {
		return "", errors.NewNotFoundError("no user found with the given credentials")
	}

	token, tokenErr := jwt.NewService("secret-Key", 3600).GenerateToken(account.Id, account.Role)
	if tokenErr != nil {
		return "", errors.NewInternalServerError("something went wrong")
	}
	return token, nil
}

func (s *service) RegisterUser(username string, email string, password string, role string) *errors.Rest {
	id, err := uuidv4.NewService().Generate()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	account, accountErr := New(id, username, email, password)
	if accountErr != nil {
		return accountErr
	}
	account.SetRole(role)
	account.CreatedOn = time.Current()

	searchResult, searchErr := s.repository.FindByEmail(email)
	if searchErr != nil && searchErr.Status != http.StatusNotFound {
		return searchErr
	}
	if searchResult != nil {
		return errors.NewBadRequestError("email is already registered")
	}

	if dbErr := s.repository.Save(account); dbErr != nil {
		return dbErr
	}
	return nil
}


func (s *service) Search(id string) (*Account, *errors.Rest) {
	account, err := s.repository.Find(id)
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (s *service) Delete(id string) *errors.Rest {
	_, err := s.Search(id)
	if err != nil {
		return err
	}

	if err := s.repository.Delete(id); err != nil {
		return err
	}
	return nil
}
