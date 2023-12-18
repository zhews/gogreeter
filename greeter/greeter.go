package greeter

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"go.opentelemetry.io/otel"
)

var (
	genericCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gogreeter_greetings_generic_total",
		Help: "Greetings without a name specified.",
	})
	specificCounter = promauto.NewCounter(prometheus.CounterOpts{
		Name: "gogreeter_greetings_specific_total",
		Help: "Greetings with a name specified.",
	})
	tracer = otel.GetTracerProvider().Tracer("Greeter")
)

func Greet(ctx context.Context, name string) string {
	cleanedName := strings.TrimSpace(name)
	if len(cleanedName) == 0 {
		genericCounter.Inc()
		return "Hello World!"
	}
	validateName(ctx, cleanedName)
	specificCounter.Inc()
	return fmt.Sprintf("Hello %s!", name)
}

func validateName(ctx context.Context, _ string) {
	_, span := tracer.Start(ctx, "validate-name")
	defer span.End()
	seconds := time.Second * time.Duration(rand.Intn(4))
	time.Sleep(seconds)
}
