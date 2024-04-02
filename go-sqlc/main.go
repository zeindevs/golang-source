package main

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/zeindevs/gosqlc/products"
)

func main() {
	connStr := "postgres://postgres:root@localhost:5432/sqlctest?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx := context.Background()
	queries := products.New(db)

	// create product
	product, err := queries.CreateProduct(ctx, products.CreateProductParams{
		Name:      "Coffee Machine",
		Price:     "199.99",
		Available: sql.NullBool{Bool: true, Valid: true},
	})
	if err != nil {
		log.Fatal(err)
	}

	// list all products
	productList, err := queries.ListProducts(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, element := range productList {
		log.Println(element)
	}

	prices, err := queries.TotalPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(prices)

	// get single product
	product, err = queries.GetProduct(ctx, 2)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No product with ID 1")
			return
		}
		log.Fatal(err)
	}
	log.Println(product)

	err = queries.DeleteProduct(ctx, 2)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No product with ID 2")
			return
		}

		log.Fatal(err)
	}
	log.Print("Product with ID 2 deleted")
}
