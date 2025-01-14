package models

import "time"

type Carts struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	ProductID int       `json:"product_id"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
}
