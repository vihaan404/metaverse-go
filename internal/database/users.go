// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

const createUser = `-- name: CreateUser :one

INSERT INTO users (id, username, password, role) VALUES (
  gen_random_uuid(),
  $1,
  $2,
  $3
) RETURNING id, username, password, avatar_id, role
`

type CreateUserParams struct {
	Username string   `json:"username"`
	Password password `json:"password"`
	Role     UserRole `json:"role"`
}
type password struct {
	plainText *string
	hash      string
}

func (p *password) Set(plainText string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	p.plainText = &plainText
	p.hash = hash
	return nil
}

func (p *password) Matches(plainTextPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(plainTextPassword))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Password, arg.Role)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Password,
		&i.AvatarID,
		&i.Role,
	)
	return i, err
}
