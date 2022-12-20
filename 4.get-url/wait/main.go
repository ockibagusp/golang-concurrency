package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("waiting...")
	time.Sleep(5 * time.Second)
	log.Println("wait: responding")
	io.WriteString(w, "Hello Word!\n")
}

func main() {
	http.HandleFunc("/", getRoot)
	log.Println("started wait...")

	err := http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
