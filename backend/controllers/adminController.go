package controllers

import "github.com/gofiber/fiber/v2"

// Dummy users for example
var users = []map[string]string{
	{"id": "1", "name": "Ivan"},
	{"id": "2", "name": "Nimesha"},
}

// GetStats returns dashboard statistics
func GetStats(c *fiber.Ctx) error {
	stats := map[string]interface{}{
		"totalUsers":  len(users),
		"activeUsers": 1, // Dummy number
	}
	return c.JSON(stats)
}

// GetUsers returns a mock list of users
func GetUsers(c *fiber.Ctx) error {
	return c.JSON(users)
}
