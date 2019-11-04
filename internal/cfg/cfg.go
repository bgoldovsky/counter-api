package config

import (
	"os"
	"strconv"
)

// GetExpires returns counter expiration in seconds
func GetExpires() int {

	tmp := os.Getenv("EXPIRES")
	exp, err := strconv.ParseInt(tmp, 10, 0)
	if err != nil {
		exp = 60
	}

	return int(exp)
}

// GetPort returns HTTP port
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	return port
}

// GetStore returns storage file path
func GetStore() string {
	store := os.Getenv("STORE_PATH")
	if store == "" {
		store = "./store.gob"
	}

	return store
}
