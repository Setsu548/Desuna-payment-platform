package routes

import (
	"go/token"

	"github.com/Setsu548/Desuna-payment-platform/tree/master/bank-simulator-backend/db/sqlc"
	"github.com/Setsu548/Desuna-payment-platform/tree/master/bank-simulator-backend/db/util"
)

type ServerConfig struct {
	config     util.Config
	store      sqlc.Store
	tokenMaker token.IMaker
}

func NewServer(config util.Config, store sqlc.Store) (*ServerConfig, error) {
	tokenMaker, err := token
}
