package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format(time.RFC3339)
		fmt.Fprint(w, currentTime)
	})

	fmt.Println("Server starting on port 8000...")
	http.ListenAndServe(":8000", nil)
}
