package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID   int64 `bun:",pk,autoincrement"`
	Name string
}

type Story struct {
	ID       int64
	Title    string
	AuthorID int64
	Author   *User `bun:"rel:belongs-to,join:author_id=id"`
}

func ConnectDB() *bun.DB {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))
	return db
}

func tableRelationsips() {
	db := ConnectDB()

	story := new(Story)
	err := db.NewSelect().Model(story).Relation("Author").Limit(1).Scan(context.TODO())
	if err != nil {
		panic(err)
	}
}

func scanningQueryResults() {
	db := ConnectDB()

	user := new(User)
	err := db.NewSelect().Model(user).Limit(1).Scan(context.TODO())
	if err != nil {
		panic(err)
	}

	// into scalars
	var id int64
	var name string
	err = db.NewSelect().Model((*User)(nil)).Column("id", "name").Limit(1).Scan(context.TODO(), &id, &name)
	if err != nil {
		panic(err)
	}

	// into a map[string]interface{}
	var m map[string]interface{}
	err = db.NewSelect().Model((*User)(nil)).Column("id", "name").Limit(1).Scan(context.TODO(), &m)
	if err != nil {
		panic(err)
	}

	// into slices of the types
	var users []User
	err = db.NewSelect().Model(&users).Limit(1).Scan(context.TODO())
	if err != nil {
		panic(err)
	}

	var ids []int64
	var names []string
	err = db.NewSelect().Model((*User)(nil)).Column("id", "name").Limit(1).Scan(context.TODO(), &ids, &names)
	if err != nil {
		panic(err)
	}

	var ms []map[string]interface{}
	err = db.NewSelect().Model((*User)(nil)).Scan(context.TODO(), &ms)
	if err != nil {
		panic(err)
	}

	_, err = db.NewDelete().Model((*User)(nil)).Returning("id").Exec(context.TODO(), &ids)
	if err != nil {
		panic(err)
	}
}

func definingModels() {
	db := ConnectDB()

	// Create users table.
	_, err := db.NewCreateTable().Model((*User)(nil)).Exec(context.TODO())
	if err != nil {
		panic(err)
	}

	// Drop users table.
	_, err = db.NewDropTable().Model((*User)(nil)).Exec(context.TODO())
	if err != nil {
		panic(err)
	}

	// Drop and create tables.
	err = db.ResetModel(context.TODO(), (*User)(nil))
	if err != nil {
		panic(err)
	}

	// Insert a single user.
	user := &User{Name: "admin"}
	_, err = db.NewInsert().Model(user).Exec(context.TODO())
	if err != nil {
		panic(err)
	}

	// Insert multiple users (bulk-insert).
	users := []User{{Name: "user1"}, {Name: "user2"}}
	_, err = db.NewInsert().Model(&users).Exec(context.TODO())
	if err != nil {
		panic(err)
	}

	// Update rows
	userUpdate := &User{ID: 1, Name: "admin"}
	_, err = db.NewUpdate().Model(userUpdate).Column("name").WherePK().Exec(context.TODO())
	if err != nil {
		panic(err)
	}

	// Delete rows
	userDelete := &User{ID: 1}
	_, err = db.NewDelete().Model(userDelete).WherePK().Exec(context.TODO())
	if err != nil {
		panic(err)
	}

	// Select rows scanning the results
	userSelect := new(User)
	err = db.NewSelect().Model(userSelect).Where("id = ?", 1).Scan(context.TODO())
	if err != nil {
		panic(err)
	}

	// Select first 10 users.
	var usersAll []User
	err = db.NewSelect().Model(&usersAll).OrderExpr("id ASC").Limit(10).Scan(context.TODO())
	if err != nil {
		panic(err)
	}
}

func existingode() {
	db := ConnectDB()

	users := make([]User, 0)

	err := db.NewRaw("SELECT id, name FROM ? LIMIT ?", bun.Ident("users"), 100).Scan(context.TODO(), &users)
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}

	if _, err := tx.Exec("...existing query"); err != nil {
		panic(err)
	}

	var model User
	_, err = db.NewInsert().Conn(tx).Model(&model).Exec(context.TODO())
}

func connecting() {
	db := ConnectDB()
	// res, err := db.ExecContext(context.TODO(), "SELECT 1")

	var num int
	err := db.QueryRowContext(context.TODO(), "SELECT 1").Scan(&num)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result 1: %d\n", num)

	// res, err := db.NewSelect().ColumnExpr("1").Exec(context.TODO())

	err = db.NewSelect().ColumnExpr("1").Scan(context.TODO(), &num)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Result 2: %d\n", num)
}
