package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/plain")
		_, err := w.Write([]byte("Hello world!"))	
		if err != nil {
			log.Printf("http write error: %v\n", err)
		}
	})

	fmt.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
