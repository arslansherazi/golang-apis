package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string
	Price    int
	Platform string
	Tags     []string
	Ratings  float64
}

func main() {
	courses := []Course{
		{"Artificial Intelligence", 1950, "Udemy", []string{"ai", "a-z", "udemy"}, 4.2},
		{"Go Lang", 3500, "Coursera", []string{"go-lang", "speed", "udemy"}, 4.2},
	}

	finalJson, error := json.MarshalIndent(courses, "", "\t")
	if error != nil {
		panic(error)
	}
	fmt.Printf("Type of json data:: %T", finalJson)
	fmt.Println()
	fmt.Printf("%s\n", finalJson)
}
