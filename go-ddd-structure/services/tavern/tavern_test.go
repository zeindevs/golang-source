package tavern_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/zeindevs/tavern/domain/product"
	"github.com/zeindevs/tavern/services/order"
	"github.com/zeindevs/tavern/services/tavern"
)

func init_products(t *testing.T) []product.Product {
	beer, err := product.NewProduct("Beer", "Healthy Beverage", 1.90)
	if err != nil {
		t.Fatal(err)
	}

	peenuts, err := product.NewProduct("Peanuts", "Snacks", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	wine, err := product.NewProduct("Wine", "nasty drink", 0.99)
	if err != nil {
		t.Fatal(err)
	}

	return []product.Product{beer, peenuts, wine}
}

func TestTavern(t *testing.T) {
	products := init_products(t)

	os, err := order.NewOrderService(
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	tavern, err := tavern.NewTavern()
	if err != nil {
		t.Fatal(err)
	}

	uid, err := os.AddCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		t.Fatal(err)
	}
}
