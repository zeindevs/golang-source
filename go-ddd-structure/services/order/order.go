package order

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/zeindevs/tavern/domain/customer"
	cstMemory "github.com/zeindevs/tavern/domain/customer/memory"
	"github.com/zeindevs/tavern/domain/customer/mongo"
	"github.com/zeindevs/tavern/domain/product"
	pdMemory "github.com/zeindevs/tavern/domain/product/memory"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customers customer.Repository
	products  product.Repository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	o := &OrderService{}
	// Loop through all the cfgs and apply them
	for _, cfg := range cfgs {
		if err := cfg(o); err != nil {
			return nil, err
		}
	}
	return o, nil
}

func WithCustomerRepository(cr customer.Repository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cstMemory.New()
		return nil
	}
}

func WithMongoCustomerRepository(ctx context.Context, connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, connectionString)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}

func WithMemoryProductRepository(products []product.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := pdMemory.New()
		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func (os *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	// Fetch the customer
	c, err := os.customers.Get(customerID)
	if err != nil {
		return 0, err
	}
	// Get each Product
	var products []product.Product
	var total float64

	for _, id := range productsIDs {
		p, err := os.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("Customer: %s has ordered %d products", c.GetID(), len(products))
	return total, nil
}

func (os *OrderService) AddCustomer(name string) (uuid.UUID, error) {
	c, err := customer.NewCustomer(name)
	if err != nil {
		return uuid.Nil, err
	}

	if err := os.customers.Add(c); err != nil {
		return uuid.Nil, err
	}

	return c.GetID(), nil
}
