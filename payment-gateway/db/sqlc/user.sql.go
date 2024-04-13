package sqlc

import (
	"context"

	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/model"
)

const createUsers = `-- name: CreateUsers :one
INSERT INTO merchant.users (
    username,
    hashed_password,
    full_name,
    email
) VALUES (
    $1, $2, $3, $4
) RETURNING username, hashed_password, full_name, email, password_changed_at, created_at
`

type CreateUsersParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUsersParams) (model.User, error) {
	row := q.db.QueryRowContext(ctx, createUsers,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)

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
