package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
)

func main() {
    http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
    })

	http.Handle("/fourohfour", http.NotFoundHandler())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostname, err := os.Hostname()
		if err != nil {
			panic(err)
		}

		w.Header().Set("content-type", "text/plain")
		greeting := fmt.Sprintf("Hello world! I run on %s", hostname)
		_, err = w.Write([]byte(greeting))	
		if err != nil {
			log.Printf("http write error: %v\n", err)
		}
	})

	fmt.Println("Starting...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

