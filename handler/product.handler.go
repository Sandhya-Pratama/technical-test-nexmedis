package handler

import (
	"strconv"

	"github.com/Sandhya-Pratama/technical-test-nexmedis/models"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/repository"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/service"
	"github.com/gofiber/fiber/v2"
)

// Initialize product service
var productService = &service.ProductService{
	ProductRepo: &repository.ProductRepository{},
}

// CreateProductHandler handles the creation of a new product
func CreateProductHandler(c *fiber.Ctx) error {
	var product models.Product
	if err := c.BodyParser(&product); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Call service to create product
	err := productService.CreateProduct(&product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Product created successfully",
		"product": product,
	})
}

// GetProductHandler retrieves a product by ID
func GetProductHandler(c *fiber.Ctx) error {
	// Ambil ID dari parameter URL
	idStr := c.Params("id")

	// Konversi ID dari string ke int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid product ID"})
	}

	// Panggil service untuk mendapatkan produk berdasarkan ID
	product, err := productService.GetProduct(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Product not found"})
	}
	return c.JSON(product)
}

// SearchProductHandler searches products by name or category
func SearchProductHandler(c *fiber.Ctx) error {
	name := c.Query("name")
	category := c.Query("category")
	products, err := productService.SearchProducts(name, category)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Search failed"})
	}
	return c.JSON(products)
}
