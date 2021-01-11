package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("world")
	})
	test := app.Group("/test")
	{
		test.Get("/hello", func(ctx *fiber.Ctx) error {
			return ctx.JSON("test")
		})
	}

	app.Listen(":3000")
}
