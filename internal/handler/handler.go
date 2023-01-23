package handler

import "github.com/gofiber/fiber/v2"

func InitRoutes() *fiber.App {
	router := fiber.New()

	router.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	library := router.Group("/library")
	library.Get("/", GetLibrary)

	return router
}
