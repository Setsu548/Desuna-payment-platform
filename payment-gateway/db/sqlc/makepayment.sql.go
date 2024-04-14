package sqlc

import "context"

type MakePaymentParams struct {
}

// MakePayment implements Querier.
func (q *Queries) MakePayment(ctx context.Context) (Payment, error) {
	panic("unimplemented")
}
