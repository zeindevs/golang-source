package main

import (
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}

	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}

	fmt.Println("new account => ", acc.Number)

	return acc
}

func seedAccounts(s Storage) {
	seedAccount(s, "zee", "dev", "hunter9999")
}
