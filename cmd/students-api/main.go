package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Prateek2593/students-api/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//database setup
	//set up router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to students-api"))
	})
	//setup server
	server := http.Server{
		Addr:    cfg.HTTPServer.Addr,
		Handler: router,
	}

	fmt.Printf("Server started %s", cfg.HTTPServer.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start server")
	}
}
