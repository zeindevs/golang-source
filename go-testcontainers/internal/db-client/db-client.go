package db_client

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DbClient struct {
	ConnString string
	Db         *sql.DB
}

type Profile struct {
	Name string
}

func NewDbClient(connStr string) (*DbClient, error) {
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfullt pinged to the database")
	return &DbClient{
		ConnString: connStr,
		Db:         db,
	}, nil
}
