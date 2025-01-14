package router

import (
	"github.com/Sandhya-Pratama/technical-test-nexmedis/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {

	app.Post("users/register", handler.RegisterUser)
	app.Post("users/login", handler.LoginUser)

	app.Post("/products", handler.CreateProductHandler)
	app.Get("/products/:id", handler.GetProductHandler)
	app.Get("/products", handler.SearchProductHandler)

	app.Post("/carts", handler.CreateCart)
	app.Get("/carts/:id", handler.GetCartByUserID)
	app.Delete("/carts/:id", handler.DeleteCart)
}
