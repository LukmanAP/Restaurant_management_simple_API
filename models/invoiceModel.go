package models

import "time"

type Invoice struct {
	Id               string    `json:"id"`
	Invoice_id       string    `json:"invoice_id"`
	Order_id         string    `json:"order_id"`
	Payment_method   string    `json:"payment_method" validate:"eq=CARD|eq=CASH|eq=ONLINE"`
	Payment_status   string    `json:"payment_status" validate:"eq=PENDING|eq=COMPLETED|eq=FAILED"`
	Payment_due_date string    `json:"payment_due_date"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
