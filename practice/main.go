package main

import (
	"fmt"
)

func main() {

	var value = 22
	var p = &value
	fmt.Println("Value of p :", *p)
	fmt.Println("Value of p :", p)

	*p = *p / 2
	fmt.Println("Value of p :", *p)
	fmt.Println("Value of p :", p)
	fmt.Println("Value of p :", value)

}
