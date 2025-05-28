package user

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	DbPool *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepository {
	return &UserRepository{DbPool: db}
}

func (r *UserRepository) CreateUser(name, email, hash string) error {
	_, err := r.DbPool.Exec(context.Background(),
		`INSERT INTO users (name, email, password_hash) VALUES ($1, $2, $3)`,
		name, email, hash)
	return err
}

func (r *UserRepository) FindByEmail(email string) (*User, error) {
	row := r.DbPool.QueryRow(context.Background(),
		`SELECT id, name, email, password_hash FROM users WHERE email = $1`, email)

	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.PasswordHash)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
