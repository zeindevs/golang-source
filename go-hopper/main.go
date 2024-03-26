package main

import (
	"fmt"
	"log"

	"github.com/zeindevs/hopper/hopper"
)

func main() {
	user := map[string]string{
		"name": "Anthony",
		"age":  "36",
	}
	db, err := hopper.New()
	if err != nil {
		log.Fatal(err)
	}
	// _, err = db.CreateCollection("users")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	id, err := db.Insert("users", user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", id)
}
