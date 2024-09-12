package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("😀", 5)
	want := "😀😀😀😀😀"

	if got != want {
		t.Errorf("Got %q but want %q", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	result := Repeat("😎", 3)
	fmt.Println(result)
	// Output: 😎😎😎
}
