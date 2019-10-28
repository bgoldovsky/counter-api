package main

import (
	"log"

	api "github.com/bgoldovsky/counter-api/api"
	cfg "github.com/bgoldovsky/counter-api/internal/cfg"
)

func main() {

	port := cfg.GetPort()
	store := cfg.GetStore()
	exp := cfg.GetExpires()

	log.Printf("init config..\nport: %v\nstore: %v\nexpires: %v\n", port, store, exp)

	api.StartServer(port, store, exp)
}
