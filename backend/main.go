package main

import (
	"admin-dashboard/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Optional: Root route to test backend is running
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Backend is running")
	})

	// Admin dashboard routes
	routes.RegisterAdminRoutes(app)

	app.Listen(":3000")
}
