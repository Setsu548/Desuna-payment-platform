package controller

import (
	"net/http"

	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/db/sqlc"
	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/model"
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

	// send info to the bank
	// validate card number function // todo

	//vaidate ccv number function // todo

	arg := sqlc.MakePaymentParams{
		User:        req.User,
		Amount:      req.Amount,
		PaymentType: req.PaymentType,
		Email:       req.Email,
		CardNumber:  req.CardNumber,
		CardName:    req.CardName,
		ExpireDate:  req.ExpireDate,
	}

	err := sc.Store.MakePayment(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusAccepted, "ok")
}
