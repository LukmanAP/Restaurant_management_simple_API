package models

import "time"

type Note struct {
	ID        string    `json:"id"`
	Text      string    `json:"text"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Note_id   string    `json:"note_id"`
}
