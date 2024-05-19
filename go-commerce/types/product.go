package types

import (
	"fmt"
	"strings"
)

const (
	minProductNameLen = 3
)

type Product struct {
	ID   string `bson:"_id,omitempty" json:"id"`
	SKU  string `bson:"sku" json:"sku"`
	Name string `bson:"name" json:"name"`
	Slug string `bson:"slug" json:"slug"`
}

type ProductCreateRequest struct {
	SKU  string `json:"sku"`
	Name string `json:"name"`
}

func NewProductFormRequest(req *ProductCreateRequest) (*Product, error) {
	if err := ValidateProductCreateRequest(req); err != nil {
		return nil, err
	}

	parts := strings.Split(strings.ToLower(req.Name), " ")
	slug := strings.Join(parts, "-")

	return &Product{
		SKU:  req.SKU,
		Name: req.Name,
		Slug: slug,
	}, nil
}

func ValidateProductCreateRequest(req *ProductCreateRequest) error {
	if len(req.SKU) < 3 {
		return fmt.Errorf("the SKU of the product is to short")
	}

	if len(req.Name) < minProductNameLen {
		return fmt.Errorf("the name of the product is to short")
	}

	return nil
}
