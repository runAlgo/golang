package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Student struct {
}

func successHandler(w http.ResponseWriter, r *http.Request) {

	// How to set headers
	w.Header().Set("Content-Type", "application/json")
	// How to set Status-Code
	w.WriteHeader(http.StatusOK)

	res := map[string]any{
		"ok":       true,
		"message":  "JSON encode successfull",
		"datetime": time.Now().UTC(),
	}

	_ = json.NewEncoder(w).Encode(res)

}
func main() {
	http.HandleFunc("/ok", successHandler)
	fmt.Println("Joining port:3000")
	err := http.ListenAndServe(":3000", nil)
	fmt.Println(err)
}
