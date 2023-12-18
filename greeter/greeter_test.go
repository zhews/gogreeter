package greeter

import (
	"context"
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	testCases := []struct {
		name           string
		expectedOutput string
	}{
		{
			name:           "zhews",
			expectedOutput: "Hello zhews!",
		},
		{
			name:           "gogreeter",
			expectedOutput: "Hello gogreeter!",
		},
		{
			name:           "",
			expectedOutput: "Hello World!",
		},
		{
			name:           " ",
			expectedOutput: "Hello World!",
		},
	}
	ctx := context.Background()
	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Name: %s", testCase.name), func(tc *testing.T) {
			output := Greet(ctx, testCase.name)
			if testCase.expectedOutput != output {
				tc.Fatalf("expected: %s, actual: %s", testCase.expectedOutput, output)
			}
		})
	}
}
