package handler

import (
	"github.com/Sandhya-Pratama/technical-test-nexmedis/models"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/repository"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/service"
	"github.com/gofiber/fiber/v2"
)

var cartService = &service.CartService{
	CartRepo: &repository.CartRepository{},
}

func GetCartByUserID(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	cart, err := cartService.GetCartByUserID(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get cart"})
	}
	return c.Status(fiber.StatusOK).JSON(cart)
}

func CreateCart(c *fiber.Ctx) error {
	var carts models.Carts
	if err := c.BodyParser(&carts); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	createdCart, err := cartService.CreateCart(&carts)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create cart"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Cart created successfully",
		"data":    createdCart,
	})
}

func DeleteCart(c *fiber.Ctx) error {
	userID, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	err = cartService.DeleteCart(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to delete cart"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Cart deleted successfully"})
}
