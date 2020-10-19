package fibonacci

import "testing"

func TestCalculate(t *testing.T) {
	got := Calculate(10)
	want := 55
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}
