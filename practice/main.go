package main

import (
	"fmt"
	"time"
)

func main() {

	// Here's a basic switch

	i := 3
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	// You can use commans to seperate multiple expressions in the same case statement

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Yayy, It's the Weekend")
	default:
		fmt.Println("Nah, It's just a Weekday")
	}
	//switch without an expression is an alternate way to express if/else logic

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// A type switch compares types instead of values. Could be used to discover the type of an interface value
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a ini")
		default:
			fmt.Printf("Don't know what type %T I am\n", t)

		}
	}
	whatAmI(true)
	whatAmI(5)
	whatAmI("string ? or what !!")
}
