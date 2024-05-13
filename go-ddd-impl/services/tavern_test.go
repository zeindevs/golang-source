package services

import (
	"context"
	"testing"

	"ddd-impl/aggregate"

	"github.com/google/uuid"
)

func TestTavern(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		// WithMemoryCustomerRepository(),
		WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017/tavern"),
		WithMemoryProductRepository(products),
	)

	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	cust, err := aggregate.NewCustomer("percy")
	if err != nil {
		t.Fatal(err)
	}

	if err := os.customers.Add(cust); err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Fatal(err)
	}
}
