package iteration

import "testing"

// Benchmarking
// The testing.B gives you access to the cryptically named b.N
// When the benchmark code is executed, it runs b.N times and measures how long it takes
// To benchmark, run go test -bench=.
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("exxpected %q, but got %q", expected, repeated)
	}
}