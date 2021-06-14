package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	dow := rand.Intn(7) + 1
	fmt.Println("Day", dow)

	var result string

	switch dow {
	case 1:
		result = "Wow ! It's Sunday o/"
	case 2:
		result = "It's Monday, Go to Sleep"
	default:
		result = "It's some other day !!"
	}

	fmt.Println(result)

}
