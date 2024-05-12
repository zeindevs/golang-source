package product

import (
	"errors"

	"github.com/google/uuid"
	"github.com/zeindevs/tavern"
)

var (
	ErrMissingValues = errors.New("missing important value")
)

type Product struct {
	item     *tavern.Item
	price    float64
	quantity int
}

func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrMissingValues
	}

	return Product{
		item: &tavern.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 100,
	}, nil
}

func (p *Product) GetID() uuid.UUID {
	return p.item.ID
}

func (p *Product) GetPrice() float64 {
	return p.price
}

func (p *Product) GetQuantity() int {
	return p.quantity
}
