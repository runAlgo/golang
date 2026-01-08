package main

import (
	"fmt"
	"net/http"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only Get is allowed", http.StatusMethodNotAllowed)
		return
	}
	_, _ = w.Write([]byte("Hello from GO net/http server"))
}

func main() {
	http.HandleFunc("/hello", handleHello)
	fmt.Println("try going to 3000 port")
	err := http.ListenAndServe(":3000", nil)
	fmt.Println(err)
}
