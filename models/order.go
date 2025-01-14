// models/order.go
package models

import "time"

type Order struct {
	ID         int       `json:"id"`
	CustomerID int       `json:"customer_id"`
	ProductID  int       `json:"product_id"`
	OrderDate  time.Time `json:"order_date"`
	Amount     float64   `json:"amount"`
}
