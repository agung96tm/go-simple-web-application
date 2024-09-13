package data

import (
	"context"
	"database/sql"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (m *UserModel) GetAll() (*[]User, error) {
	stmt := `SELECT id, name FROM users ORDER BY id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := m.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &users, nil
}

func (m *UserModel) GetById(id string) (*User, error) {
	stmt := `SELECT id, name FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var user User
	err := m.DB.QueryRowContext(ctx, stmt, id).Scan(&user.ID, &user.Name)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *UserModel) Create(user *User) error {
	stmt := `INSERT INTO users (name) VALUES ($1) RETURNING id`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, stmt, user.Name).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (m *UserModel) Update(user *User) error {
	stmt := `UPDATE users SET name = $1 WHERE id = $2`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, stmt, user.Name, user.ID)
	if err != nil {
		return err
	}

	return nil
}

func (m *UserModel) Delete(id string) error {
	stmt := `DELETE FROM users WHERE id = $1`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := m.DB.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}
	return nil
}
