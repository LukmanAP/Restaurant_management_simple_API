package models

import "time"

type Food struct {
	ID         string    `json:"id"`
	Name       string    `json:"name" validate:"required, min=2, max=100"`
	Price      int       `json:"price" validate:"required"`
	Food_image string    `json:"food_image" validate:"required"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Food_id    string    `json:"food_id"`
	Menu_id    string    `json:"menu_id" validate:"required"`
}
