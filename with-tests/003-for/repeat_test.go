package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	testEquals(t, Repeat("a", 5), "aaaaa")
	testEquals(t, Repeat("b", 1), "b")
	testEquals(t, Repeat("", 10), "")
}
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	var s string = Repeat("a", 5)
	fmt.Print(s)
	//Output: aaaaa
}

// Helpers
func testEquals(t testing.TB, got, want string) {
	t.Helper()
	if want != got {
		t.Errorf("expected %q but got %q", want, got)
	}
}
