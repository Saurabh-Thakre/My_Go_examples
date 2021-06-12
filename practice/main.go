package main

import (
	"bufio"
	"fmt"
	"os"
)

const aConst int = 65

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Entered String :", input)

}
