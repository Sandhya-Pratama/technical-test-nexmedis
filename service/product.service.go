package service

import (
	"time"

	"github.com/Sandhya-Pratama/technical-test-nexmedis/models"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/repository"
)

type ProductService struct {
	ProductRepo *repository.ProductRepository
}

// CreateProduct creates a new product
func (service *ProductService) CreateProduct(product *models.Product) error {
	product.CreatedAt = time.Now() // Set CreatedAt to the current time
	return service.ProductRepo.CreateProduct(product)
}

// GetProduct retrieves a product by its ID
func (service *ProductService) GetProduct(id int) (*models.Product, error) {
	return service.ProductRepo.GetProductByID(id)
}

// SearchProducts searches products by name and category
func (service *ProductService) SearchProducts(name string, category string) ([]models.Product, error) {
	return service.ProductRepo.SearchProducts(name, category)
}
