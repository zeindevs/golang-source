package main

import (
	"fmt"
	"net/http"
)

type DB interface {
	Store(string) error
}

type Store struct{}

func (s *Store) Store(value string) error {
	fmt.Println("storing into db", value)
	return nil
}

func myExecuteFunc(db DB) ExecuteFn {
	return func(s string) {
		// access to DB??
		fmt.Println("my ex func", s)
		db.Store(s)
	}
}

// This is comming from a third party lib
type ExecuteFn func(string)

func Execute(fn ExecuteFn) {
	fn("FOO BAR BAZ")
}

func makeHTTPFunc(db DB, fn httpFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(db, w, r); err != nil {
			//
		}
	}
}

type httpFunc func(db DB, w http.ResponseWriter, r *http.Request) error

func handler(db DB, w http.ResponseWriter, r *http.Request) error {
	return nil
}

func main() {
	s := &Store{}
	http.HandleFunc("/", makeHTTPFunc(s, handler))
	Execute(myExecuteFunc(s))
}
