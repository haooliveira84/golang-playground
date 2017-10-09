package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Person struct {
	State    string `json:"State"`
	City     string `json:"City"`
	Fullname string `json:"FullName"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing parameter")
		return
	}
	csvFile, err := os.Open(os.Args[1])
	file := csv.NewReader(bufio.NewReader(csvFile))
	if err != nil {
		fmt.Println("Can't read file")
		panic(err)
	}
	var people []Person
	for {
		row, error := file.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		people = append(people, Person{
			State:    row[0],
			City:     row[1],
			Fullname: row[2],
		})
	}
	peopleJson, _ := json.Marshal(people)
	fmt.Println(string(peopleJson))
}
