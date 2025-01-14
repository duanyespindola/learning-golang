package main

import "fmt"

// Animal interface defines the behavior
type Animal interface {
	Speak() string
}

// GenericAnimal struct serves as the base implementation
type GenericAnimal struct {
	Name  string
	Sound string
}

// Speak method for GenericAnimal
func (a GenericAnimal) Speak() string {
	return fmt.Sprintf("The %s does %s", a.Name, a.Sound)
}

// Dog struct embeds GenericAnimal
type Dog struct {
	GenericAnimal
}

// Cat struct embeds GenericAnimal
type Cat struct {
	GenericAnimal
}

func main() {
	a := 1
	var animal Animal

	if a == 1 {
		animal = Dog{GenericAnimal{Name: "dog", Sound: "woff woff"}}
	} else {
		animal = Cat{GenericAnimal{Name: "cat", Sound: "meaw"}}
	}

	fmt.Println(animal.Speak())
}
