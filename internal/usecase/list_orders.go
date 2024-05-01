package usecase

import (
	"github.com/allanmaral/go-expert-clean-arch-challenge/internal/entity"
)

type ListOrdersOutput struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() ([]ListOrdersOutput, error) {
	orders, err := c.OrderRepository.GetAll()
	if err != nil {
		return nil, err
	}

	output := make([]ListOrdersOutput, len(orders))
	for i, order := range orders {
		output[i] = ListOrdersOutput{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.Price + order.Tax,
		}
	}

	return output, nil
}
