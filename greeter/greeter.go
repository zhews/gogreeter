package greeter

import (
	"fmt"
	"strings"
)

func Greet(name string) string {
	cleanedName := strings.TrimSpace(name)
	if len(cleanedName) == 0 {
		return "Hello World!"
	}
	return fmt.Sprintf("Hello %s!", name)
}
