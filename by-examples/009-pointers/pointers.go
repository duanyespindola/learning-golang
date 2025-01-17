package main

import "fmt"

type V struct {
	x int
	y int
}

func main() {
	a := V{x: 1, y: 2}
	b := a

	fmt.Println(`a`, a, `b`, b)

	b.x = 3

	fmt.Println(`a`, a, `b`, b)

	c := &a
	c.y = 4

	fmt.Println(`a`, a, `b`, b, `c`, c)

	//.
	//.

	d := &V{5, 6}
	e := d
	fmt.Println(`d`, d, `e`, e)

	d.x = 7
	e.y = 8
	fmt.Println(`d`, d, `e`, e)

}
