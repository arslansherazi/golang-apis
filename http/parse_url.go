package main

import (
	"fmt"
	"net/url"
)

func main() {
	const employeesUrl string = "http://localhost:8000/api/v1/employees?emp_id=23&emp_code=EMP23"

	// parse url
	result, _ := url.Parse(employeesUrl)

	fmt.Println("Protocol::", result.Scheme)
	fmt.Println("Host::", result.Host)
	fmt.Println("Path::", result.Path)
	fmt.Println("Port::", result.Port())
	fmt.Println("Raw Query::", result.RawQuery)

	queryParams := result.Query()

	fmt.Println("Query Params::")
	for queryParamName, queryParamValue := range queryParams {
		fmt.Printf("%s:: \t %s", queryParamName, queryParamValue[0])
		fmt.Println()
	}
}
