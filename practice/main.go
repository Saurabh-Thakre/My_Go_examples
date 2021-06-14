package main

import (
	"fmt"
)

func main() {

	states := make(map[string]string)
	fmt.Println(states)
	states["WA"] = "Washington"
	states["CA"] = "California"
	states["NO"] = "New Orleans"
	fmt.Println(states)

	orleans := states["NO"]
	fmt.Println(orleans)

}
