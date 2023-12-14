package main

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zhews/gogreeter/greeter"
)

func main() {
	app := fiber.New()

	app.Get("/", getGreeting)
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	const ADDRESS = ":8080"
	if err := app.Listen(ADDRESS); err != nil {
		log.Fatal("failed to listen on address", err)
	}
}

func getGreeting(c *fiber.Ctx) error {
	input := c.Query("name")
	output := greeter.Greet(input)
	return c.Status(http.StatusOK).SendString(output)
}
