package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 11, "Woof Woof"}
	fmt.Printf("%+v\n", poodle)
	poodle.Speak()
	poodle.Sound = "Arf"
	poodle.Speak()

}

type Dog struct {
	Breed  string
	Weight int
	Sound  string
}

func (d Dog) Speak() {

	fmt.Println(d.Sound)
}
