package services

import (
	"github.com/Setsu548/Desuna-payment-platform/tree/master/payment-gateway/model"
	"github.com/go-playground/validator/v10"
)

var ValidCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if Currency, ok := fl.Field().Interface().(model.CurrencyType); ok {
		return Currency.IsValid()
	}
	return false
}
