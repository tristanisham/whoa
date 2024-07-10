package main

import (
	"log"
	"net/http"

	"github.com/dunglas/frankenphp"
)

func main() {
	if err := frankenphp.Init(); err != nil {
		panic(err)
	}
	defer frankenphp.Shutdown()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		req, err := frankenphp.NewRequestWithContext(r, frankenphp.WithRequestDocumentRoot("./public", false))
		if err != nil {
			panic(err)
		}

		if err := frankenphp.ServeHTTP(w, req); err != nil {
			panic(err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
