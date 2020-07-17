package uuidv4

import (
	"errors"
	"github.com/google/uuid"
)

type service struct {
}

func NewService() Service{
	return &service{}
}

func (s *service) Generate() (string, error) {
	uuid4, err := uuid.NewRandom()
	if err != nil {
		return "", errors.New("error while generating uuid")
	}
	uuidStr := uuid4.String()
	if uuidStr == "" {
		return "", errors.New("error while generating uuid")
	}
	return uuidStr, nil
}

func (s *service) Validate(id string) error {
	_, err := uuid.Parse(id)
	if err != nil {
		return errors.New("invalid uuid")
	}
	return nil
}