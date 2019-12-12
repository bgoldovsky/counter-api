package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bgoldovsky/counter-api/internal/api"
	"github.com/bgoldovsky/counter-api/internal/cfg"
	"github.com/bgoldovsky/counter-api/internal/dal"
	"github.com/bgoldovsky/counter-api/internal/services"
)

func main() {
	port := cfg.GetPort()
	path := cfg.GetPath()
	expires := cfg.GetExpires()

	log.Printf("init config..\nport: %v\nstore: %v\nexpires: %v\n", port, path, expires)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	repo, err := dal.New(path)
	if err != nil {
		log.Fatalln("repository creation error", err)
	}

	service, err := services.NewCounterService(repo, expires)
	if err != nil {
		log.Fatalln("service creation error", err)
	}

	server := api.NewServer(service)
	server.StartServer(port)

	log.Println("application stopped")
}
