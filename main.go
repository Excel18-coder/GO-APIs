package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name   string `json:"username"`
	Age    int    `json:"age"`
	Active bool   `json:"is_active"`
}

func main() {
	p := Person{
		Name:   "Excel",
		Age:    20,
		Active: true,
	}

	f, err := os.Create("output.json")
	if err != nil {
		fmt.Println("error creating file", err)
		panic(err)
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(p)
	if err != nil {
		fmt.Println("error encoding person: ", err)
		panic(err)
	}

}
