package cfg

import (
	"os"
	"strconv"
)

// GetExpires returns counter expiration in seconds
func GetExpires() int {

	tmp := os.Getenv("EXPIRES")
	expires, err := strconv.ParseInt(tmp, 10, 0)
	if err != nil {
		expires = 60
	}

	return int(expires)
}

// GetPort returns HTTP port
func GetPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}

	return port
}

// GetPath returns storage file path
func GetPath() string {
	path := os.Getenv("STORE_PATH")
	if path == "" {
		path = "./store.gob"
	}

	return path
}
