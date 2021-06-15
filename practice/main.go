package main

import (
	"fmt"
)

func main() {

	myfunction()
	sum := sumvalues(2, 5)
	fmt.Println("Sum is: ", sum)
}

func myfunction() {
	fmt.Println("This is my function")
}

func sumvalues(value1 int, value2 int) int {
	return value1 + value2
}
