package main

import (
	"fmt"
)

//Structure

func main() {
	poodle := Dog{"Poodle", 10}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)

}

// Here Dog is a struct
type Dog struct {
	Breed  string
	Weight int
}
