package repository

import (
	"fmt"

	"github.com/Sandhya-Pratama/technical-test-nexmedis/db"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/models"
)

type ProductRepository struct {
}

// CreateProduct inserts a new product into the database
func (repo *ProductRepository) CreateProduct(product *models.Product) error {
	query := `INSERT INTO products (name, description, price, stock, category, created_at) 
              VALUES ($1, $2, $3, $4, $5, NOW()) RETURNING id`

	err := db.DB.QueryRow(query, product.Name, product.Description, product.Price,
		product.Stock, product.Category).Scan(&product.ID)

	if err != nil {
		return fmt.Errorf("could not create product: %v", err)
	}
	return nil
}

// GetProductByID retrieves a product by its ID
func (repo *ProductRepository) GetProductByID(id int) (*models.Product, error) {
	var product models.Product
	query := `SELECT id, name, description, price, stock, category, created_at FROM products WHERE id = $1`

	err := db.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description,
		&product.Price, &product.Stock, &product.Category, &product.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("could not get product: %v", err)
	}
	return &product, nil
}

// SearchProducts retrieves a list of products matching search criteria
func (repo *ProductRepository) SearchProducts(name string, category string) ([]models.Product, error) {
	query := `SELECT id, name, description, price, stock, category, created_at FROM products 
              WHERE name ILIKE $1 AND category ILIKE $2`

	rows, err := db.DB.Query(query, "%"+name+"%", "%"+category+"%")
	if err != nil {
		return nil, fmt.Errorf("could not search products: %v", err)
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var product models.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price,
			&product.Stock, &product.Category, &product.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("could not scan product: %v", err)
		}
		products = append(products, product)
	}
	return products, nil
}
