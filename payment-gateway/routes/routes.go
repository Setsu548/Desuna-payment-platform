package routes

import (
	"fmt"

	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/db/sqlc"
	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/db/util"
	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/services"
	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/token"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type ServerConfig struct {
	config     util.Config
	store      sqlc.Store
	tokenMaker token.IMaker
	router     *gin.Engine
}

func NewServer(config util.Config, store sqlc.Store) (*ServerConfig, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &ServerConfig{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("Currency", services.ValidCurrency)
		if err != nil {
			return nil, nil
		}
	}

	server.setupRouter()

	return server, nil
}

func (server *ServerConfig) setupRouter() {
	route := gin.Default()

	route.GET("/route", func(ctx *gin.Context) {})
}

func (server *ServerConfig) Start(address string) error {
	return server.router.Run(address)
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
