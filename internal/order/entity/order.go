package entity

import (
	"errors"
	"math"
	"github.com/google/uuid"
)

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(price float64, tax float64) (*Order, error) {
	id := uuid.New().String()
	order := Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := order.IsValid()
	if err != nil {
		return nil, err
	}

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

func (o *Order) CalculateFinalPrice() error {
	err := o.IsValid()
	if err != nil {
		return err
	}
	o.FinalPrice = o.Price * (1 + o.Tax/100)
	return nil
}
