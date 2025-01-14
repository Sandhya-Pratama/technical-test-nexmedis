package service

import (
	"github.com/Sandhya-Pratama/technical-test-nexmedis/models"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/repository"
)

type CartService struct {
	CartRepo *repository.CartRepository
}

func (s *CartService) CreateCart(cart *models.Carts) (models.Carts, error) {
	return s.CartRepo.CreateCart(cart)
}

func (s *CartService) GetCartByUserID(userID int) (*models.Carts, error) {
	return s.CartRepo.GetCartByUserID(userID)
}

func (s *CartService) DeleteCart(userID int) error {
	return s.CartRepo.DeleteCart(userID)
}
