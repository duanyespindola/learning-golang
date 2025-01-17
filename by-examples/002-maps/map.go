package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"k1": 1, "k2": 2}
	a, b := m["k3"]
	fmt.Println(a, b)
}
