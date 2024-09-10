package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	db *sql.DB
}

type PostgresConfig struct {
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSSLMode  string
}

func NewPostgresStorage(cfg PostgresConfig) *PostgresStorage {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Connected")

	return &PostgresStorage{db: db}
}

func (s *PostgresStorage) Init() (*sql.DB, error) {

	if err := s.createUsersTable(); err != nil {
		return nil, err
	}

	if err := s.createBooksTable(); err != nil {
		return nil, err
	}

	if err := s.createHightlightsTable(); err != nil {
		return nil, err
	}

	return s.db, nil
}

func (s *PostgresStorage) createUsersTable() error {
	_, err := s.db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
      id SERIAL,
      email VARCHAR(255) NOT NULL,
      firstName VARCHAR(255) NOT NULL,
      lastName VARCHAR(255) NULL,
      createdAt TIMESTAMP NOT NULL DEFAULT NOW(),
      PRIMARY KEY (id),
      UNIQUE (email)
    );
  `)
	return err
}

func (s *PostgresStorage) createBooksTable() error {
	_, err := s.db.Exec(`
    CREATE TABLE IF NOT EXISTS books (
      isbn VARCHAR(255) NOT NULL,
      title VARCHAR(255) NOT NULL,
      authors VARCHAR(255) NOT NULL,
      createdAt TIMESTAMP NOT NULL DEFAULT NOW(),

      PRIMARY KEY (isbn),
      UNIQUE (isbn)
    );
  `)
	return err
}

func (s *PostgresStorage) createHightlightsTable() error {
	_, err := s.db.Exec(`
    CREATE TABLE IF NOT EXISTS highlights (
      id SERIAL,
      text TEXT,
      location VARCHAR(255) NOT NULL,
      note TEXT,
      userId BIGINT NOT NULL,
      bookID VARCHAR(255) NOT NULL,
      createdAt TIMESTAMP NOT NULL DEFAULT NOW(),

      PRIMARY KEY (id),
      FOREIGN KEY (userId) REFERENCES users(id)
    );
  `)
	return err
}
