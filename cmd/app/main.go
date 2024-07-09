package main

import (
	"go-proxy/internal/config"
	"go-proxy/internal/router"
	"log"
	"net/http"
)

func main() {
	cfg := config.MustLoad()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	r := router.NewRouter()

	addr := ":" + cfg.HttpServer.Port
	log.Printf("Starting server on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
