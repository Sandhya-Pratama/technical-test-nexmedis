package handler

import (
	"fmt"

	"github.com/Sandhya-Pratama/technical-test-nexmedis/models"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/repository"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/service"
	"github.com/gofiber/fiber/v2"
)

var userService = &service.UserService{
	UserRepo: &repository.UserRepository{},
}

func RegisterUser(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	// Menambahkan log untuk melihat request body yang diterima
	fmt.Println("Request Body:", user)

	err := userService.RegisterUser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User registered successfully",
		"user":    fiber.Map{"username": user.Username, "email": user.Email},
	})
}

func LoginUser(c *fiber.Ctx) error {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request payload"})
	}

	user, err := userService.LoginUser(req.Username, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"message": "Login successful",
		"user":    fiber.Map{"username": user.Username, "email": user.Email},
	})
}
