package repository

import (
	"github.com/Sandhya-Pratama/technical-test-nexmedis/db"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/models"
)

type OrderRepository struct{}

func (r *OrderRepository) CreateOrder(order *models.Order) (int, error) {
	query := "INSERT INTO orders (customer_id, product_id, order_date, amount) VALUES ($1, $2, $3, $4) RETURNING id"
	var id int
	err := db.DB.QueryRow(query, order.CustomerID, order.ProductID, order.OrderDate, order.Amount).Scan(&id)
	return id, err
}

func (r *OrderRepository) GetOrderByID(id int) (*models.Order, error) {
	query := "SELECT id, customer_id, product_id, order_date, amount FROM orders WHERE id = $1"
	var order models.Order
	err := db.DB.QueryRow(query, id).Scan(&order.ID, &order.CustomerID, &order.ProductID, &order.OrderDate, &order.Amount)
	return &order, err
}
