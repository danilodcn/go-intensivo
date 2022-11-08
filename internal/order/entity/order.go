package entity

import (
	"errors"
	"math"
)

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price float64, tax float64) (*Order, error) {
	order := Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.IsValid()
	if err != nil {
		return nil, err
	}
	order.FinalPrice = order.CalculateFinalPrice()
	return &order, nil
}

func (o *Order) IsValid() error {
	if o.ID == "" {
		return errors.New("invalid ID")
	}
	if o.Price <= 0 {
		return errors.New("invalid Price")
	}
	if math.Abs(o.Tax) > 100 {
		return errors.New("invalid Tax")
	}
	return nil
}

func (o *Order) CalculateFinalPrice() float64 {
	FinalPrice := o.Price * (1 + o.Tax/100)
	return FinalPrice
}
