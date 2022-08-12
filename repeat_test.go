package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertEqual := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("expected '%q' but got '%q'", want, got)
		}
	}
	got := Repeat("a", 5)
	want := "aaaaa"
	assertEqual(t, got, want)
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
func ExampleRepeat() {
	x := Repeat("x", 9)
	fmt.Println(x)
	//Output: xxxxxxxxx
}
