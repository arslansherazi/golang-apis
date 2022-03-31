package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	const getUserUrl string = "https://reqres.in/api/users/2"

	response, error := http.Get(getUserUrl)
	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	fmt.Println("Status Code::", response.StatusCode)
	fmt.Println("Response Length::", response.ContentLength)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Response::")
	fmt.Println(string(content))
}
