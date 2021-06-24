package main

import "fmt"

func main() {

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
		//Above is the basic one with a single condition
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
		// This is the classic initial/condition/after for loop
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
