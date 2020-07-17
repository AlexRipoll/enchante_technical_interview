package user

import (
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/uuidv4"
	"github.com/AlexRipoll/market/pkg/time"
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
	// TODO implementation
	return "", nil
}

func (s *service) Search(id string) (*Account, *errors.Rest) {
	// TODO implementation
	return nil, nil
}

func (s *service) Delete(id string) *errors.Rest {
	// TODO implementation
	return nil
}

func (s *service) FindAll() ([]Account, *errors.Rest) {
	// TODO implementation
	return nil, nil
}
