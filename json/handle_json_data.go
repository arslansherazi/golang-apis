package main

import (
	"encoding/json"
	"fmt"
)

type Course struct {
	Name     string   `json:"course_name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Tags     []string `json:"tags,omitempty"`
	Ratings  float64  `json:"ratings"`
}

func createJson() []byte {
	courses := []Course{
		{"Artificial Intelligence", 1950, "Udemy", []string{"ai", "a-z", "udemy"}, 4.2},
		{"Go Lang", 3500, "Coursera", []string{"go-lang", "speed", "udemy"}, 4.2},
	}

	finalJson, error := json.MarshalIndent(courses, "", "\t")
	if error != nil {
		panic(error)
	}

	fmt.Printf("Type of json data:: %T", finalJson)
	fmt.Printf("\n\n")

	return finalJson
}

func consumeJson(jsonData []byte) {
	if json.Valid(jsonData) {
		var courses []Course
		json.Unmarshal(jsonData, &courses)
		for index := range courses {
			course := courses[index]
			fmt.Println("Name::", course.Name)
			fmt.Println("Price::", course.Price)
			fmt.Println("Platform::", course.Platform)
			fmt.Println("Tags::", course.Tags)
			fmt.Println("Ratings::", course.Ratings)
			fmt.Println()
			fmt.Println()
		}
	} else {
		fmt.Println("Json is not valid")
	}
}

func main() {
	finalJson := createJson()
	consumeJson(finalJson)
}
