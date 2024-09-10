package main

import (
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
}

type Storage interface {
	CreateBook(Book) error
	CreateHighlights([]Highlight) error
	GetBookByISBN(string) (*Book, error)
	GetActiveUsers() ([]*User, error)
	GetRandomHightlights(int) ([]*Highlight, error)
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) CreateBook(b Book) error {
	_, err := s.db.Exec(`
    INSERT INTO books (isbn, title, authors)
    VALUES ($1, $2, $3)
  `, b.ISBN, b.Title, b.Authors)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) CreateHighlights(hs []Highlight) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := `INSERT INTO Highlights (text, location, note, userId, bookId) VALUES `
	values := []interface{}{}
	for _, h := range hs {
		query += "($1, $2, $3, $4, $5),"
		values = append(values, h.Text, h.Location, h.Note, h.UserID, h.BookID)
	}

	query = query[:len(query)-1]

	_, err = s.db.Exec(query, values...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetBookByISBN(isbn string) (*Book, error) {
	rows, err := s.db.Query(`
    SELECT * FROM books WHERE isbn = $1 LIMIT 1
  `, isbn)
	if err != nil {
		return nil, err
	}
	book := new(Book)
	for rows.Next() {
		if err := rows.Scan(&book.ISBN, &book.Title, &book.Authors); err != nil {
			return nil, err
		}
	}

	if book.ISBN == "" {
		return nil, fmt.Errorf("book not found")
	}

	return book, nil
}

func (s *Store) GetActiveUsers() ([]*User, error) {
	return nil, nil
}

func (s *Store) GetRandomHightlights(total int) ([]*Highlight, error) {
	return nil, nil
}
