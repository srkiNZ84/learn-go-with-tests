package iteration

import (
	"fmt"
	"testing"
)

func TestRepeated(t *testing.T) {
	t.Run("Print a five times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

	t.Run("Print b thirteen times", func(t *testing.T) {
		got := Repeat("b", 13)
		want := "bbbbbbbbbbbbb"

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}

func ExampleRepeat() {
	r := Repeat("Z", 20)
	fmt.Println(r)
	// Output: ZZZZZZZZZZZZZZZZZZZZ
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("s", 20)
	}
}
