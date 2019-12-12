package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	srv "github.com/bgoldovsky/counter-api/internal/services"
)

var counter srv.Counter

type CounterServer struct {
	handler http.Handler
}

func NewServer() *CounterServer {
	return &CounterServer{}
}

// StartServer func starts listening single endpoint of service
func (s *CounterServer) StartServer(port string, filepath string, expires int) {
	log.Printf("service starting on port %s\n", port)

	c, err := srv.NewCounter(filepath, expires)
	if err != nil {
		log.Fatal("can't create counter")
		return
	}

	counter = c

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), counterHandler{}))
}

// ServeHTTP handler
func (h counterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Println(r.Method, r.URL.Path)

		err := counter.Increment()
		if err != nil {
			writeError(w, err)

			return
		}

		state := counter.GetState()
		val, err := json.Marshal(state)
		if err != nil {
			writeError(w, err)

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(string(val)))
	}
}

func writeError(w http.ResponseWriter, err error) {
	const text = "internal server Error"
	log.Println(text, err.Error())

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(text))
}
