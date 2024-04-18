package controller

import (
	"net/http"

	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/db/sqlc"
	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/model"
	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/services"
	"github.com/gin-gonic/gin"
)

type PaymentsRequest struct {
	User        string  `json:"user"`
	Amount      float32 `json:"amount"`
	PaymentType string  `json:"payment_type"` // visa/mastercard/other
	Email       string  `json:"email"`
	CardNumber  int64   `json:"card_number"`
	CardName    string  `json:"card_name"`
	CCV         string  `json:"ccv_number"`
	ExpireDate  string  `json:"expire_date"`
}

type PaymentResponse struct{}

func NewPaymentResponse(paymen model.Payment) PaymentResponse {
	return PaymentResponse{}
}

func (sc *ServerConfig) MakePayment(ctx *gin.Context) {
	var req PaymentsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, ErrorResponse(err))
		return
	}

	// validate card number function // todo
	//vaidate ccv number function // todo

	// send info to the bank
	bankRequest := PaymentsRequest{
		User:        req.User,
		Amount:      req.Amount,
		PaymentType: req.PaymentType,
		Email:       req.Email,
		CardNumber:  req.CardNumber,
		CardName:    req.CardName,
		CCV:         req.CCV,
		ExpireDate:  req.ExpireDate,
	}

	// it have to receive from Bank service
	_, err := services.BankService(ctx, bankRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	arg := sqlc.MakePaymentParams{
		User:        req.User,
		Amount:      req.Amount,
		PaymentType: req.PaymentType,
		Email:       req.Email,
		CardNumber:  req.CardNumber,
		CardName:    req.CardName,
		ExpireDate:  req.ExpireDate,
	}

	err = sc.Store.MakePayment(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, map[string]string{"status": "Created"})
}
