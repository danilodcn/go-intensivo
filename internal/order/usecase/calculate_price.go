package usecase

import (
	"github.com/danilodcn/go-intensivo/internal/order/entity"
	"github.com/danilodcn/go-intensivo/internal/order/infra/database"
)

type OrderInputDTO struct {
	ID    string `default:""`
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPriceUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPriceUseCase(orderRepository *database.OrderRepository) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *CalculateFinalPriceUseCase) Execute(input *OrderInputDTO) (*OrderOutputDTO, error) {
	order, err := entity.NewOrder(input.Price, input.Tax)
	if input.ID != "" {
		order.ID = input.ID
	}

	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrice()

	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
