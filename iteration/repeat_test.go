package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		expected := "aaaaa"
		if got != expected {
			t.Errorf("expected '%q' | got '%q'", got, expected)
		}
	})
}

// To run the benchmark: go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	// Runs the function b.N times
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
