package bl

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	// Test Println
	Println("Hello, World!")

	// Test Printf
	Printf("Hello, %s!", "World")

	// Test Print
	Print("Hello, World without newline")

	Printf("Hello %d", 42)
	Println("Error %w", fmt.Errorf("an error occurred"))
}
