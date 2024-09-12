package iteration

import "testing"

func TestIterate(t *testing.T) {
	got := Iterate("😀")
	want := "😀😀😀😀😀"

	if got != want {
		t.Errorf("Got %q but want %q", got, want)
	}
}
