package main

import "fmt"

func main() {
	var mystring = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
	var sample = []byte(mystring)
	// var sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	for i := 0; i < len(sample); i++ {
		fmt.Printf("%q ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)
}
