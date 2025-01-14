package models

import "time"

type Order struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
}
