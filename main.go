package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/zhews/gogreeter/greeter"
)

func main() {
	app := fiber.New()

	app.Get("/", getGreeting)

	const ADDRESS = ":8080"
	app.Listen(ADDRESS)
}

func getGreeting(c *fiber.Ctx) error {
	input := c.Query("name")
	output := greeter.Greet(input)
	return c.Status(http.StatusOK).SendString(output)
}
