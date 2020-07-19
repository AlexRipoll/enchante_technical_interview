package product

import (
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/time"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/uuidv4"
)

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{repository}
}

func (s *service) Add(name string, price float64, sellerId string) *errors.Rest {
	uuidService := uuidv4.NewService()
	id, err := uuidService.Generate()
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	if err := uuidService.Validate(sellerId); err != nil {
		return errors.NewBadRequestError(err.Error())
	}

	product, productErr := New(id, name, price, sellerId)
	if productErr != nil {
		return productErr
	}
	product.CreatedOn = time.Current()

	if err := s.repository.Save(product); err != nil {
		return err
	}
	return nil
}
func (s *service) Find(id string) (*Product, *errors.Rest) {
	p, err := s.repository.Find(id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (s *service) Update(id string, name string, price float64,  sellerId string) *errors.Rest {
	_, err := s.Find(id)
	if err != nil {
		return err
	}

	p, err := New(id, name, price, sellerId)
	if err != nil {
		return err
	}
	p.UpdatedOn = time.Current()

	if err = s.repository.Update(p); err != nil {
		return err
	}
	return nil
}

func (s *service) Delete(id string,  sellerId string) *errors.Rest {
	return nil
}

func (s *service) FindAll() ([]Product, *errors.Rest) {
	return nil, nil
}
//
//func (s *service) Search(id string) (*Account, *errors.Rest) {
//	account, err := s.repository.Find(id)
//	if err != nil {
//		return nil, err
//	}
//	return account, nil
//}
//
//func (s *service) Delete(id string) *errors.Rest {
//	_, err := s.Search(id)
//	if err != nil {
//		return err
//	}
//
//	if err := s.repository.Delete(id); err != nil {
//		return err
//	}
//	return nil
//}
