package product

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	"github.com/zeindevs/go-ecommerce/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}

	products := make([]*types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func (s *Store) GetProductByID(id int) (*types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products WHERE id = $1 LIMIT 1", id)
	if err != nil {
		return nil, err
	}

	p := new(types.Product)
	for rows.Next() {
		p, err = scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
	}

	if p.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return p, nil
}

func (s *Store) GetProductsByIDs(ids []int) ([]types.Product, error) {
	idsStr := strings.Join(idsToString(ids), ",")
	rows, err := s.db.Query("SELECT * FROM products WHERE id IN (" + idsStr + ")")
	if err != nil {
		return nil, err
	}

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := scanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) CreateProduct(p *types.Product) error {
	stmt, err := s.db.Prepare("INSERT INTO products (name, description, image, price, quantity) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		p.Name,
		p.Description,
		p.Image,
		p.Price,
		p.Quantity,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateProduct(p types.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = $1, price = $2, image = $3, description = $4, quantity = $5 WHERE id = $6", p.Name, p.Price, p.Image, p.Description, p.Quantity, p.ID)
	if err != nil {
		return err
	}
	return nil
}

func scanRowsIntoProduct(rows *sql.Rows) (*types.Product, error) {
	p := new(types.Product)
	err := rows.Scan(
		&p.ID,
		&p.Name,
		&p.Description,
		&p.Image,
		&p.Price,
		&p.Quantity,
		&p.CreatedAt,
	)
	return p, err
}

func idsToString(ids []int) []string {
	idsStr := make([]string, len(ids))
	for i, v := range ids {
		idsStr[i] = strconv.Itoa(v)
	}
	return idsStr
}
