// Package main is the main entry of the app
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Tecu23/ws-server/pkg/server"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 8088, "Server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialize a new logger which writes to std out, prefixed with curr date & tim
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	app.serve()
}

func (app *application) serve() error {
	server := server.NewServer()

	server.SetupServer()

	fmt.Println("Server started at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}

	return nil
}
