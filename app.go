package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var filename = string(os.Args[1])
	if len(filename) < 2 {
		fmt.Println("Missing parameter")
	}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Can't read file", filename)
		panic(err)
	}
	fmt.Println(string(file))
}
