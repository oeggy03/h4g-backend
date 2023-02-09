package controller

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// must "remove" the cookie by creating another cookie that has already expired
func SignOut(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour), //expired an hour ago
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Logout success.",
	})
}
