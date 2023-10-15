package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("repeat 5 times", func(t *testing.T) {
		actual := Repeat("a", 3)
		expected := "aaa"

		if actual != expected {
			t.Errorf("expected %q but got %q", expected, actual)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
