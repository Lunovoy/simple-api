package handler

import "github.com/gofiber/fiber/v2"

func GetLibrary(c *fiber.Ctx) error {
	return c.SendString("From GetLibrary")
}
