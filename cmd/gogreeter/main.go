package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gofiber/contrib/otelfiber/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/zhews/gogreeter/greeter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

func main() {
	ctx := context.Background()
	tracerProvider := setupOpenTelemetryTracerProvider(ctx)
	defer func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			log.Fatal("could not shutdown tracer provider", err)
		}

	}()
	otel.SetTracerProvider(tracerProvider)

	app := fiber.New()

	app.Use(otelfiber.Middleware())

	app.Get("/", getGreeting)
	app.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler()))

	const ADDRESS = ":8080"
	if err := app.Listen(ADDRESS); err != nil {
		log.Fatal("failed to listen on address", err)
	}
}

func getGreeting(c *fiber.Ctx) error {
	input := c.Query("name")
	output := greeter.Greet(c.UserContext(), input)
	return c.Status(http.StatusOK).SendString(output)
}

func setupOpenTelemetryTracerProvider(ctx context.Context) *trace.TracerProvider {
	resource, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("gogreeter"),
		),
	)
	if err != nil {
		log.Fatal("could not setup OpenTelemetry", err)
	}
	exporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithEndpoint("jaeger:4317"), otlptracegrpc.WithInsecure())
	if err != nil {
		log.Fatal("could not setup OpenTelemetry", err)
	}
	return trace.NewTracerProvider(
		trace.WithResource(resource),
		trace.WithBatcher(exporter),
	)
}
