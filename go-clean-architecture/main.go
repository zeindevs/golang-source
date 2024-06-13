package main

import (
	"clean-architecture/modules/categories"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PORT = ":3000"

func handler(http http.ResponseWriter, r *http.Request) {

}

func main() {
	db, err := gorm.Open(postgres.Open(""), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	categoryRepository := categories.NewCategoryRepository(db)
	_ = categoryRepository

	r := mux.NewRouter()
	r.HandleFunc("/login", handler).Methods("GET")
	r.HandleFunc("/products", handler).Methods("POST")
	r.HandleFunc("/products", handler).Methods("GET")
	r.HandleFunc("/products/{id}", handler).Methods("GET")
	r.HandleFunc("/products/{id}", handler).Methods("PUT")
	r.HandleFunc("/products/{id}", handler).Methods("DELETE")

	r.HandleFunc("/categories", handler).Methods("POST")
	r.HandleFunc("/categories", handler).Methods("GET")
	r.HandleFunc("/categories/{id}", handler).Methods("GET")
	r.HandleFunc("/categories/{id}", handler).Methods("PUT")
	r.HandleFunc("/products/{id}", handler).Methods("DELETE")

	fmt.Println("starting web server at localhost", PORT)
	if err := http.ListenAndServe(PORT, r); err != nil {
		log.Fatal(err)
	}
}
