package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthCheck(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func createCpuFunc(loopCnt int) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		for i := 0; i < loopCnt; i++ {
		}
		log.Printf("cpu")
		fmt.Fprintf(w, "hello\n")
	}
}
