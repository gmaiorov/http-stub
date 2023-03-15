package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Got %s request to %s\n", r.Method, r.RequestURI)
	_, _ = io.WriteString(w, "0xDEADBEEF\n")
}

func main() {
	fmt.Println("Starting http stub server...")

	http.HandleFunc("/", handle)
	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Println("Stopping server...")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
