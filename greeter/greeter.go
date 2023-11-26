package greeter

import "fmt"

func Greet(name string) string {
	if len(name) == 0 {
		return "Hello World!"
	}
	return fmt.Sprintf("Hello %s!", name)
}
