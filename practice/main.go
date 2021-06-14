package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Enter Correct Password To Enter !")
	input1, _ := reader.ReadString('\n')
	float2, err := strconv.ParseInt(strings.TrimSpace(input1), 0, 32)

	if float2 == 42 {

		fmt.Println("You've granted the permission, you can enter !")
	} else if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Intruder Detected!!!, Calling Security")
	}
}
