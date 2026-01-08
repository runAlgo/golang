package main

import (
	"fmt"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "this method not allowed", http.StatusMethodNotAllowed)
	}
	w.Write([]byte("Welcome to Dashboard try to /hello?name=kalu"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "this method not allowed", http.StatusMethodNotAllowed)
	}

	name := r.URL.Query().Get("name")

	if name == "" {
		name = "Guest"
	}


	_, _ = w.Write([]byte(name))
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/hello", helloHandler)
	fmt.Println("Going to port: 3000")
	err := http.ListenAndServe(":3000", nil)
	fmt.Println(err)
}
