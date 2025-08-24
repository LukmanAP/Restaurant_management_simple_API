package models

import "time"

type Menu struct {
	ID         string    `json:"id"`
	Name       string    `json:"name" validate:"required"`
	Category   string    `json:"category" validate:"required"`
	Start_date time.Time `json:"start_date"`
	End_date   time.Time `json:"end_date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Menu_id    string    `json:"menu_id"`
}
