package main

import (
	"context"
	"log"
	"net/http"

	"github.com/zeindevs/gocommerce/api"
	"github.com/zeindevs/gocommerce/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	router := http.NewServeMux()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	productStore := store.NewMongoProductStore(client.Database("gocommerce"))
	productHandler := api.NewProductHandler(productStore)

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Bro"))
	})
	router.HandleFunc("GET /product/{id}", productHandler.HandleGetProductByID())
	router.HandleFunc("GET /product", productHandler.HandleGetAllProduct())
	router.HandleFunc("POST /product", productHandler.HandlePostProduct())

	app := http.Server{
		Addr:    ":3001",
		Handler: router,
	}

	log.Println("app listening on 0.0.0.0::3001")
	app.ListenAndServe()
}
