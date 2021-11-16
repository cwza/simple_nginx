package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthCheck(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func simple(w http.ResponseWriter, req *http.Request) {
	log.Printf("simple")
	fmt.Fprintf(w, "hello\n")
}
