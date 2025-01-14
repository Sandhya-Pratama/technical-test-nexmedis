package service

import (
	"github.com/Sandhya-Pratama/technical-test-nexmedis/models"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/repository"
)

type OrderService struct {
	OrderRepo *repository.OrderRepository
}

func (s *OrderService) CreateOrder(order *models.Order) (int, error) {
	return s.OrderRepo.CreateOrder(order)
}

func (s *OrderService) GetOrderByID(id int) (*models.Order, error) {
	return s.OrderRepo.GetOrderByID(id)
}
