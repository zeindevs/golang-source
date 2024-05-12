package memory

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/zeindevs/tavern/domain/product"
)

type MemoryRepository struct {
	products map[uuid.UUID]product.Product
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		products: make(map[uuid.UUID]product.Product),
	}
}

func (mr *MemoryRepository) GetAll() ([]product.Product, error) {
	products := []product.Product{}
	mr.Lock()
	for _, p := range mr.products {
		products = append(products, p)
	}
	mr.Unlock()
	return products, nil
}

func (mr *MemoryRepository) GetByID(uid uuid.UUID) (product.Product, error) {
	var p product.Product
	mr.Lock()
	if _, ok := mr.products[uid]; !ok {
		return product.Product{}, fmt.Errorf("product not found :%w", product.ErrProductNotFound)
	}
	p = mr.products[uid]
	mr.Unlock()
	return p, nil
}

func (mr *MemoryRepository) Add(p product.Product) error {
	if mr.products == nil {
		mr.Lock()
		mr.products = make(map[uuid.UUID]product.Product)
		mr.Unlock()
	}
	// Meke sure customer is already in repo
	if _, ok := mr.products[p.GetID()]; ok {
		return fmt.Errorf("product already exists :%w", product.ErrProductAlreadyExists)
	}
	mr.Lock()
	mr.products[p.GetID()] = p
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Update(p product.Product) error {
	mr.Lock()
	if _, ok := mr.products[p.GetID()]; !ok {
		return fmt.Errorf("product not found :%w", product.ErrProductNotFound)
	}
	mr.products[p.GetID()] = p
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Delete(uid uuid.UUID) error {
	mr.Lock()
	if _, ok := mr.products[uid]; !ok {
		return fmt.Errorf("product not found :%w", product.ErrProductNotFound)
	}
	delete(mr.products, uid)
	mr.Unlock()
	return nil
}
