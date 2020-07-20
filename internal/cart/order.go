package cart

import "github.com/AlexRipoll/enchante_technical_interview/pkg/errors"

type Order struct {
	Id     string `json:"id"`
	UserId string `json:"user_id"`
	Items  []Item `json:"items"`
}

type Item struct {
	Id       string  `json:"id"`
	Price    float64 `json:"price"`
	Quantity int8    `json:"quantity"`
	SellerId string  `json:"seller_id"`
}

type Service interface {
	Purchase(userId string, items []Item) *errors.Rest
}

type Repository interface {
	Save(order *Order) *errors.Rest
}

func New(id, userId string, items []Item) (*Order, *errors.Rest) {
	o := &Order{
		Id:     id,
		UserId: userId,
		Items:  items,
	}
	return o, nil
}
