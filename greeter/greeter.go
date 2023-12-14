package greeter

import (
	"fmt"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
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
)

func Greet(name string) string {
	cleanedName := strings.TrimSpace(name)
	if len(cleanedName) == 0 {
		genericCounter.Inc()
		return "Hello World!"
	}
	specificCounter.Inc()
	return fmt.Sprintf("Hello %s!", name)
}
