package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	srv "github.com/bgoldovsky/counter-api/internal/services"
)

type counterHandler struct {
	http.Handler
}

// StartServer func starts listening single endpoint of service
func StartServer(port string, filepath string, expires int) {
	log.Printf("Service starting on port %s\n", port)

	srv.LoadState(filepath, expires)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), counterHandler{}))
}

// ServeHTTP handler
func (h counterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		log.Println(r.Method, r.URL.Path)

		state, err := srv.GetState()

		if err != nil {
			writeError(w, err)

			return
		}

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
