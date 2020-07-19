package product

import (
	"fmt"
	"github.com/AlexRipoll/enchante_technical_interview/pkg/errors"
	"html"
	"strings"
)

const (
	minNameLength = 2
	minPriceValue = 0.01
)

type Product struct {
	Id string `json:"id"`
	SellerId string `json:"seller_id"`
	Name string `json:"name"`
	Price float64 `json:"price"`
	CreatedOn string `json:"created_on"`
	UpdatedOn string `json:"updated_on"`
}

type Service interface {
	Find(id string) (*Product, *errors.Rest)
	Add(name string, price float64,  sellerId string) *errors.Rest
	Update(id string, name string, price float64,  sellerId string) *errors.Rest
	Delete(id string) *errors.Rest
	FindAll() ([]Product, *errors.Rest)
}

type Repository interface {
	Find(id string) (*Product, *errors.Rest)
	Save(product *Product) *errors.Rest
	Update(product *Product) *errors.Rest
	Delete(id string) *errors.Rest
	FindAll() ([]Product, *errors.Rest)
}

func New(id string, name string, price float64, sellerId string) (*Product, *errors.Rest) {
	p := Product{
		Id:        id,
		SellerId:  sellerId,
		Name:      name,
		Price:     price,
	}
	p.validateName()
	p.validatePrice()

	return &p, nil
}

func (p *Product) validateName() *errors.Rest {

	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	if p.Name == "" {
		return errors.NewBadRequestError("username can't be null")
	}
	if len(p.Name) < minNameLength {
		return errors.NewBadRequestError(
			fmt.Sprintf("name must be at least %d characters long", minNameLength))
	}
	return nil
}

func (p *Product) validatePrice() *errors.Rest {
	if p.Price <  minPriceValue {
		return errors.NewBadRequestError(
			fmt.Sprintf("minmum price value is %f", minPriceValue))
	}
	return nil
}