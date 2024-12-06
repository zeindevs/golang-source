package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func ConnectDB() error {
	db, err := sql.Open("sqlite3", "./todo.db")
	if err != nil {
		return err
	}
	DB = db
	log.Println("database connected")
	return nil
}

func CloseDB() error {
	return DB.Close()
}

func SetupDB() error {
	_, err := DB.Exec(`create table if not exists tasks (
  id integer not null primary key,
  title text,
  completed boolean default false,
  position integer
)`)
	if err != nil {
		return err
	}
	log.Println("create table successful")
	return nil
}
