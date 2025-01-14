package handler

import (
	"github.com/Sandhya-Pratama/technical-test-nexmedis/models"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/repository"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/service"
	"github.com/gofiber/fiber/v2"
)

var orderService = &service.OrderService{
	OrderRepo: &repository.OrderRepository{},
}

func CreateOrderHandler(c *fiber.Ctx) error {
	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		return err
	}
	id, err := orderService.CreateOrder(&order)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"id": id,
	})
}

func GetOrderHandler(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return err
	}
	order, err := orderService.GetOrderByID(id)
	if err != nil {
		return err
	}
	return c.JSON(order)
}
