package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter")
		return
	}
	var filename = string(os.Args[1])
	fmt.Println(len(filename))
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Can't read file", filename)
		panic(err)
	}
	fmt.Println(string(file))
}
