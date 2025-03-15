package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Current time is: %s", time.Now().Format(time.RFC1123))
	})
	http.ListenAndServe(":8080", nil)
}
