package controller

type PaymentsRequest struct {
	Amount         float64 `json:"amount"`
	Account        int64   `json:"account"`
	PaymentType    int     `json:"payment_type"`
	PaymentDate    *string `json:"payment_date"`
	Lang           *string `json:"lang"`
	Channel        string  `json:"channel"`
	Email          *string `json:"email"`
	PhoneNumber    *int64  `json:"phone_number"`
	PaymenMethodId int64   `json:"payment_method_id"`
}
