package main

import (
	"fmt"
	"math"
)

const min_diff = 0.000000000000001

func Sqrt(x float64) float64 {
	z := 1.0
	prev := z

	for i := 0; i < 10; i++ {
		z = z - ((z*z - x) / (2 * z))
		fmt.Printf("Step %x = %v \n", i, z)
		diff := math.Abs(prev - z)

		if diff <= min_diff {
			return prev
		} else {
			prev = z
		}
	}

	return z
}

func main() {
	fmt.Printf("My        Function = %v\n", Sqrt(122))
	fmt.Printf("math.Sqrt function = %v\n", math.Sqrt(122))
}
