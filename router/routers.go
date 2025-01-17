package router

import (
	"github.com/Sandhya-Pratama/technical-test-nexmedis/handler"
	"github.com/Sandhya-Pratama/technical-test-nexmedis/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Public Routes
	public := app.Group("/api/v1")
	public.Post("/users/register", handler.RegisterUser)
	public.Post("/users/login", handler.LoginUser)

	// Private Routes
	private := app.Group("/api/v1", middleware.AuthRequired)
	private.Post("/products", handler.CreateProductHandler)
	private.Get("/products/:id", handler.GetProductHandler)
	private.Get("/products", handler.SearchProductHandler)

	private.Post("/carts", handler.CreateCart)
	private.Get("/carts/:id", handler.GetCartByUserID)
	private.Delete("/carts/:id", handler.DeleteCart)

	private.Post("/orders", handler.CreateOrderHandler)
	private.Get("/orders/:id", handler.GetOrderHandler)
}
