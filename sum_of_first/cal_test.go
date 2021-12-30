package cal

import "testing"

func TestSumOfFirstThree(t *testing.T) {
	given := 3
	want := 6

	get := sumOfFirst(given)
	if want != get {
		t.Errorf("given %d want %d but %d\n", given, want, get)
	}
}

func TestSumOfFirstFour(t *testing.T) {
	given := 4
	want := 10

	get := sumOfFirst(given)
	if want != get {
		t.Errorf("given %d want %d but %d\n", given, want, get)
	}
}
