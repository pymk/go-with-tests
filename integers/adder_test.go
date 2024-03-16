package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("Check 2 + 2", func(t *testing.T) {
		got := Add(2, 2)
		expected := 4

		if got != expected {
			t.Errorf("expected '%d' | got '%d'", expected, got)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
