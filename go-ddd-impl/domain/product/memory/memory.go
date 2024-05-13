package memory

import (
  "sync"

  "ddd-impl/aggregate"
  "ddd-impl/domain/product"

  "github.com/google/uuid"
)

type MemoryRepository struct {
  products map[uuid.UUID]aggregate.Product
  sync.Mutex
}

func New() *MemoryRepository {
  return &MemoryRepository{
    products: make(map[uuid.UUID]aggregate.Product),
  }
}

func (mr *MemoryRepository) GetAll() ([]aggregate.Product, error) {
  var products []aggregate.Product

  for _, product := range mr.products {
    products = append(products, product)
  }

  return products, nil
}

func (mr *MemoryRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
  if product, ok := mr.products[id]; ok {
    return product, nil
  }
  return aggregate.Product{}, product.ErrProductNotFound
}

func (mr *MemoryRepository) Add(p aggregate.Product) error {
  mr.Lock()
  defer mr.Unlock()

  if _, ok := mr.products[p.GetID()]; ok {
    return product.ErrProductAlreadyExists
  }

  mr.products[p.GetID()] = p
  return nil
}

func (mr *MemoryRepository) Update(update aggregate.Product) error {
  mr.Lock()
  defer mr.Unlock()

  if _, ok := mr.products[update.GetID()]; !ok {
    return product.ErrProductNotFound
  }

  mr.products[update.GetID()] = update
  return nil
}

func (mr *MemoryRepository) Delete(id uuid.UUID) error {
  mr.Lock()
  defer mr.Unlock()

  if _, ok := mr.products[id]; !ok {
    return product.ErrProductNotFound
  }

  delete(mr.products, id)
  return nil
}
