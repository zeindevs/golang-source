package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/zeindevs/go-sqlc-generator/products"
)

func main() {
	connStr := "postgres://postgres:root@localhost:5432/sqlcgenerator?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	queries := products.New(db)

	// create product
	create, err := queries.CreateProduct(ctx, products.CreateProductParams{
		Name:      "Coffee Machine",
		Price:     "125.55",
		Available: sql.NullBool{Bool: true, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("CreateProduct %+v\n", create)

	// list all products
	productList, err := queries.ListProducts(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("ListProducts %+v\n", productList)

	// get product by id
	product, err := queries.GetProduct(ctx, create.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Product with ID", create.ID)
			return
		}
		log.Fatal(err)
	}

	fmt.Printf("GetProduct %+v\n", product)

	// total price
	total, err := queries.TotalPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("TotalPrice %+v\n", total)

	// delete product
	if err := queries.DeleteProduct(ctx, product.ID); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("DeleteProduct success")
}
