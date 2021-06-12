package main

import (
	"fmt"
)

const aConst int = 65

func main() {

	var aString string = "This is the String in Go"

	fmt.Println(aString)
	fmt.Printf("This variable's type is %T\n", aString)

	var aInteger int = 100
	fmt.Println(aInteger)

	var defaultInt int
	fmt.Println(defaultInt)

	var anotherString = "This is another String"
	fmt.Println(anotherString)
	fmt.Printf("This variable's type is %T\n", anotherString)

	var anotherInt = 45
	fmt.Println(anotherInt)
	fmt.Printf("This variable's type is %T\n", anotherInt)

	myString := "This is also a String but without var"
	fmt.Println(myString)
	fmt.Printf("This variable's type is %T\n", myString)

	fmt.Println(aConst)
	fmt.Printf("This variable's type is %T\n", aConst)

}
