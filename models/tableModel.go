package models

import "time"

type Table struct {
	ID               string    `json:"id"`
	Number_of_guests int       `json:"number_of_guests" validate:"required"`
	Table_number     int       `json:"table_number" validate:"required"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Table_id         string    `json:"table_id"`
}
