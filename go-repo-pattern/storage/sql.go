package storage

import "database/sql"

type SqlStorage struct {
	db *sql.DB
}

func NewSqlStorage(db *sql.DB) *SqlStorage {
	return &SqlStorage{
		db: db,
	}
}

func (s *SqlStorage) GetOneByID(id uint) (*User, error) {
	var user User
	query := `SELECT id, name, email FROM users WHERE id = ?`
	if err := s.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *SqlStorage) CreateUser(user *User) error {
	query := `INSERT INTO users (name, email) VALUES (?, ?)`
	_, err := s.db.Exec(query, user.Name, user.Email)
	return err
}
