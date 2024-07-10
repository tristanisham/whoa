package main

import (
	"github.com/charmbracelet/log"
	"net/http"
	"os"

	"github.com/dunglas/frankenphp"
	opts "github.com/urfave/cli/v2"
)

const VERSION = "v0.0.1"

var app = &opts.App{
	Name: "Whoa",
	Description: "Whoa is a modern PHP runtime and production webserver",
	HelpName: "whoa",
	Version: VERSION,
	Copyright: "© 2024 Tristan Isham",
	Suggest: true,
}

func main() {
	if _, ok := os.LookupEnv("WHOA_DEBUG"); ok {
		log.SetLevel(log.DebugLevel)
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}

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
