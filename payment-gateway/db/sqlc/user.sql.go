package sqlc

import (
	"context"

	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/model"
)

const createUsers = `-- name: CreateUsers :one
INSERT INTO public.users (
    username,
    hashed_password,
    full_name,
    email,
	account_number,
	bank_name
) VALUES (
    $1, $2, $3, $4, $5, $6
) RETURNING username, hashed_password, full_name, email, account_number, bank_name, password_changed_at, created_at
`

type CreateUsersParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
	AccountNumber  int64  `json:"account_number"`
	BankName       string `json:"bank_name"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUsersParams) (model.User, error) {
	row := q.db.QueryRowContext(ctx, createUsers,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
		arg.AccountNumber,
		arg.BankName,
	)

	var u model.User
	err := row.Scan(
		&u.Username,
		&u.HashedPassword,
		&u.FullName,
		&u.Email,
		&u.AccountNumber,
		&u.BankName,
		&u.PasswordChangedAt,
		&u.CreatedAt,
	)

	return u, err
}

const getUser = `-- name: GetUserByName :one
SELECT username, hashed_password, full_name, email, password_changed_at, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUserByName(ctx context.Context, username string) (model.User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var u model.User
	err := row.Scan(
		&u.Username,
		&u.HashedPassword,
		&u.FullName,
		&u.Email,
		&u.PasswordChangedAt,
		&u.CreatedAt,
	)
	return u, err
}
