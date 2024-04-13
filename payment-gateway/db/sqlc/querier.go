package sqlc

import (
	"context"
)

type Querier interface {
	MakePayment(ctx context.Context) (Payment, error)
	// ListPayment(ctx context.Context)(...)
}

var _ Querier = (*Queries)(nil)
