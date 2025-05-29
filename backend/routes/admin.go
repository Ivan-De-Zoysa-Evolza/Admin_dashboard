package routes

import (
	"admin-dashboard/controllers"

	"github.com/gofiber/fiber/v2"
)

// RegisterAdminRoutes defines admin endpoints
func RegisterAdminRoutes(app *fiber.App) {
	admin := app.Group("/admin")

	// Route for dashboard statistics
	admin.Get("/stats", controllers.GetStats)

	// Route for user list
	admin.Get("/users", controllers.GetUsers)
}
