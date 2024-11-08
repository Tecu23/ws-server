// Package main is the main entry of the app
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Tecu23/ws-server/pkg/enginepb"
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
	log.Println(app)

	// lets create a basic connection to the GRPC server
	conn, err := grpc.Dial(
		"localhost:8089",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	client := enginepb.NewChessEngineClient(conn)

	bestMove, err := client.CalculateBestMove(
		context.Background(),
		&enginepb.MoveRequest{
			Fen:  "r1bqkbnr/pppppppp/n7/8/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 1 3",
			Type: "stockfish",
		},
	)
	if err != nil {
		log.Println(err)
	}

	log.Println(bestMove)

	bestMove, err = client.CalculateBestMove(
		context.Background(),
		&enginepb.MoveRequest{
			Fen:  "r1bqkbnr/pppppppp/n7/8/4P3/5N2/PPPP1PPP/RNBQKB1R w KQkq - 1 3",
			Type: "argo",
		},
	)
	if err != nil {
		log.Println(err)
	}

	log.Println(bestMove)

	// app.serve()
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
