package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	services "github.com/bgoldovsky/counter-api/internal/services"
)

// CounterServer to handle HTTP requests
type CounterServer struct {
	http.Handler
	service services.Counter
}

// NewServer creates new server
func NewServer(service services.Counter) *CounterServer {
	return &CounterServer{service: service}
}

// StartServer func starts listening single endpoint of service
func (s *CounterServer) StartServer(port string) {
	log.Printf("service starting on port %s\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), s))
}

// ServeHTTP handler
func (s *CounterServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		log.Println(r.Method, r.URL.Path)

		err := s.service.Inc()
		if err != nil {
			s.writeError(w, err)

			return
		}

		state := s.service.State()
		val, err := json.Marshal(state)
		if err != nil {
			s.writeError(w, err)

			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(string(val)))
	}
}

func (s *CounterServer) writeError(w http.ResponseWriter, err error) {
	const text = "internal server Error"
	log.Println(text, err.Error())

	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(text))
}
