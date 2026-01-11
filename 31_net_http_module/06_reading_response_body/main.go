package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://jsonplaceholder.typicode.com/todos"

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Failed to fetch data")
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch data")
		return
	}

	fmt.Println("Data", res)
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error while Read Body")
		return
	}
	bodyText := string(bodyBytes)
	max := 300

	if (len(bodyText) < max) {
		max = len(bodyText)
	}
	fmt.Println("Length of body text:",len(bodyText))
	fmt.Println(bodyText[:max])

}
