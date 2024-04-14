package sqlc

import (
	"context"

	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/model"
)

type Querier interface {
	MakePayment(ctx context.Context) (Payment, error)
	CreateUser(ctx context.Context, arg CreateUsersParams) (model.User, error)
	// ListPayment(ctx context.Context)(...)
}

var _ Querier = (*Queries)(nil)
