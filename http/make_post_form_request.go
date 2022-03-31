package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	const usersUrl string = "https://reqres.in/api/users"

	requestData := url.Values{}

	requestData.Add("name", "Arslan Haider Sherazi")
	requestData.Add("job", "Software Engineer")

	response, error := http.PostForm(usersUrl, requestData)

	if error != nil {
		panic(error)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}
