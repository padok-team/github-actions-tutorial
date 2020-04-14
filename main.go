package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/padok-team/github-actions-tutorial/foobar"
)

const addr = ":8080"

func main() {
	http.HandleFunc("/foobar", foobarHandler)
	http.HandleFunc("/healthz", healthHandler)

	log.Printf("Listening for requests on %s\n", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// foobarHandler responds with a FooBar sequence.
// Expects a "length" parameter defining the size of the sequence.
func foobarHandler(w http.ResponseWriter, r *http.Request) {
	lengthParam := r.URL.Query().Get("length")
	if lengthParam == "" {
		http.Error(w, "missing parameter: length", http.StatusBadRequest)
		return
	}

	length, err := strconv.Atoi(lengthParam)
	if err != nil {
		http.Error(w, "invalid length", http.StatusBadRequest)
		return
	}

	seq, err := foobar.Sequence(length)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to compute sequence: %s", err.Error()), http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "%s", seq)
}

// healthHandler reports on the server's health.
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is healthy :)")
}
