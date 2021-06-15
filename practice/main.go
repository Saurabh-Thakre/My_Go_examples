package main

import (
	"fmt"
)

// for statement
func main() {

	colors := []string{"Red", "Purple", "Orange"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	// Using Range Function

	for i := range colors {
		fmt.Println(colors[i])
	}

}
