package tavern

import (
	"log"

	"github.com/google/uuid"
	"github.com/zeindevs/tavern/services/order"
)

type TavernConfiguration func(os *Tavern) error

type Tavern struct {
	// ordeservice to takes orders
	OrderService *order.OrderService

	// BillingService
	BillingService interface{}
}

func NewTavern(cfgs ...TavernConfiguration) (*Tavern, error) {
	t := &Tavern{}

	for _, cfg := range cfgs {
		if err := cfg(t); err != nil {
			return nil, err
		}
	}

	return t, nil
}

func WithOrderService(os *order.OrderService) TavernConfiguration {
	return func(t *Tavern) error {
		t.OrderService = os
		return nil
	}
}

func (t *Tavern) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := t.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}

	// Implement mongodb repository customer
	log.Printf("Bill the customer: %0.2f\n", price)
	// Billing Service

	return nil
}
