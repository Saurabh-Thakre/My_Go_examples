package main

import "fmt"

func main() {
	//Creating an int array a
	var a [5]int
	fmt.Println("emp: ", a)

	// Setting value in array for specific index
	a[4] = 100
	a[2] = 25
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	//builtin len returns the length of an array

	fmt.Println("len: ", len(a))

	// Declare and initialize an array in one line

	b := [5]int{1, 15, 234, 16, 44}
	fmt.Println("dcl :", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D :", twoD)

}
