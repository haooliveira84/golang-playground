package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var filename = string(os.Args[0])
	if len(filename) < 2 {
		fmt.Println("Missing parameter")
	}
	return
}

func openfile(filename string) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Can't read file", filename)
		panic(err)
	}
	fmt.Println(file)
}
