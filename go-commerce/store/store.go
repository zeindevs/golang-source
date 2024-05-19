package store

import (
	"context"

	"github.com/zeindevs/gocommerce/types"
)

type ProductStorer interface {
	Insert(ctx context.Context, p *types.Product) error
	GetByID(ctx context.Context, id string) (*types.Product, error)
	GetAll(ctx context.Context) ([]*types.Product, error)
	Update(ctx context.Context, p *types.Product) error
	Delete(ctx context.Context, id string) error
}
