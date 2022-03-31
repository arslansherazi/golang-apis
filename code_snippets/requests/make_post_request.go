package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	const url string = "https://reqres.in/api/users"

	requestBody := strings.NewReader(`
		{
			"name": "Arslan Haider Sherazi",
			"job": "Software Engineer"
		}
	`)

	response, error := http.Post(url, "application/json", requestBody)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
