package functional

import (
	"testing"
)

func BenchmarkToStringTakeRepeat(b *testing.B) {
	done := make(chan interface{})
	defer close(done)
	b.ResetTimer()

	for range ToString(done, Take(done, Repeat(done, "a"), b.N)) {
	}
}

func BenchmarkTakeRepeat(b *testing.B) {
	done := make(chan interface{})
	defer close(done)
	b.ResetTimer()

	for range Take(done, Repeat(done, "a"), b.N) {
	}
}
