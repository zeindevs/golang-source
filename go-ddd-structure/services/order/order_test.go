package order_test

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

func TestOrder(t *testing.T) {
	products := init_products(t)

	os, err := order.NewOrderService(
		// order.WithMongoCustomerRepository(context.Background(), "mongodb://localhost:27017"),
		order.WithMemoryCustomerRepository(),
		order.WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}

	tavern, err := tavern.NewTavern(tavern.WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}

	uid, err := os.AddCustomer("Percy")
	if err != nil {
		t.Fatal(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
		products[1].GetID(),
	}

	err = tavern.Order(uid, order)
	if err != nil {
		t.Fatal(err)
	}
}
