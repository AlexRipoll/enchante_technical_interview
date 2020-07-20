package cart

import (
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/uuidv4"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Purchase(userId string, items []Item) *errors.Rest {
	uuidService := uuidv4.NewService()
	id, err := uuidService.Generate()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	if err := uuidService.Validate(userId); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	order, orderErr := New(id, userId, items)
	if orderErr != nil {
		return orderErr
	}

	if err := s.repository.Save(order); err != nil {
		return err
	}
	return nil
}