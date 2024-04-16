package model

import "time"

type Payment struct {
	User        string    `json:"user"`
	Amount      float32   `json:"amount"`
	PaymentType string    `json:"payment_type"` // visa/mastercard/other
	Email       string    `json:"email"`
	CardNumber  int64     `json:"card_number"`
	CardName    string    `json:"card_name"`
	ExpireDate  string    `json:"expire_date"`
	CreatedAt   time.Time `json:"created_at"`
	DeletedAt   time.Time `json:"deleted_at"`
}
