package main

import (
	"fmt"
)

func giveMeTwoValues() (int, int) {
	return 10, 2
}

func main() {
	v1, v2 := giveMeTwoValues()
	fmt.Println(v1, v2)

	_, v3 := giveMeTwoValues()
	fmt.Println(v3)

	// if we try to use only one var, like this:
	//
	// v4 := giveMeTwoValues()
	// fmt.Println(v4)
	//
	// the compiler says "/multi-return-func.go:18:8: assignment mismatch: 1 variable but giveMeTwoValues returns 2 values"
	//

}
