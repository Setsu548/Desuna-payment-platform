package sqlc

import (
	"context"
)

const makePayment = `-- name: MakePayments :one
INSERT INTO public."Payments" (
	to_id_user,
	amount,
	"type",
	email,
	card_number,
	card_name,
	expire_date
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) 
-- RETURNING "amount", "type", "email", "card_number", "card_name","expire_date"
`

type MakePaymentParams struct {
	User        string  `json:"user"`
	Amount      float32 `json:"amount"`
	PaymentType string  `json:"payment_type"` // visa/mastercard/other
	Email       string  `json:"email"`
	CardNumber  int64   `json:"card_number"`
	CardName    string  `json:"card_name"`
	ExpireDate  string  `json:"expire_date"`
}

// MakePayment implements Querier.
func (q *Queries) MakePayment(ctx context.Context, arg MakePaymentParams) error {
	_, err := q.db.QueryContext(ctx, makePayment,
		arg.User,
		arg.Amount,
		arg.PaymentType,
		arg.Email,
		arg.CardNumber,
		arg.CardName,
		arg.ExpireDate,
	)

	return err
}
