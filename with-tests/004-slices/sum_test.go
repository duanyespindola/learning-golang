package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("array with fixed size", func(t *testing.T) {
		numbers := [5]int{1, 2, 3, 4, 5}
		want := 15
		got := Sum(numbers[:])
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
	// t.Run("slice with any size", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5, 6}
	// 	want := 21
	// 	got := Sum(numbers)
	// 	if got != want {
	// 		t.Errorf("got %d want %d given, %v", got, want, numbers)
	// 	}
	// })
}

func ExampleSum() {

}

func BenchmarkSum(b *testing.B) {

}
