package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Cat struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func (c Cat) convertToStruct(data string) {

}

func main() {
	url := "https://catfact.ninja/fact"

	res, err := http.Get(url)

	if err != nil {
		fmt.Println("Error fetching data")
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error while fetching data")
		return
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("read body failed", err)
		return
	}

	fmt.Println(res.StatusCode)
	fmt.Println(res.Status)

	var data Cat

	if err := json.Unmarshal(bodyBytes, &data); err != nil {
		fmt.Println("json unmarshal failed")
		return
	}

	fmt.Println("Fact:", data.Fact)
	fmt.Println("Length:", data.Length)

}
