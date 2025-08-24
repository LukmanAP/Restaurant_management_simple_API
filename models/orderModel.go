package models

import "time"

type Order struct {
	ID         string    `json:"id"`
	Order_data time.Time `json:"order_data" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Order_id   string    `json:"order_id"`
	Table_id   string    `json:"table_id" validate:"required"`
}
