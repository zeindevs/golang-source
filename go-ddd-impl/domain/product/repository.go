package product

import (
  "errors"

  "ddd-impl/aggregate"

  "github.com/google/uuid"
)

var (
  ErrProductNotFound      = errors.New("no such product")
  ErrProductAlreadyExists = errors.New("there is already such an product")
)

type ProductRepository interface {
  GetAll() ([]aggregate.Product, error)
  GetByID(uuid.UUID) (aggregate.Product, error)
  Add(aggregate.Product) error
  Update(aggregate.Product) error
  Delete(uuid.UUID) error
}
