package sqlc

import (
	"context"

	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/model"
)

type Querier interface {
	MakePayment(ctx context.Context, arg MakePaymentParams) error
	CreateUser(ctx context.Context, arg CreateUsersParams) (model.User, error)
	GetUserByName(ctx context.Context, userName string) (model.User, error)
}

var _ Querier = (*Queries)(nil)
