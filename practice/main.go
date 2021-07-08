package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	s := os.Args[1]
	res(s)

}

func res(url string) {
	req, err := http.Get(url)
	if err != nil {
		fmt.Printf("request is not hitting the end point %s", err)
	} else {
		// data, _ := ioutil.ReadAll(req.Body)
		fmt.Println(req)
		// fmt.Println(json.NewDecoder(req.Body).Decode(data))
	}

}
