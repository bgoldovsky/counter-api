package services

import (
	"log"
	"sync"
	"time"

	"github.com/bgoldovsky/counter-api/internal/dal"
	"github.com/bgoldovsky/counter-api/internal/models"
)

var counter models.Requests
var store string
var location *time.Location
var mutex sync.Mutex

func init() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		panic(err)
	}

	location = loc
}

// LoadState load current state from storage
func LoadState(path string, exp int) {
	if path == "" {
		panic("store path not specified")
	}

	store = path

	val, err := dal.Get(store)
	if err != nil {
		counter = models.Requests{
			Map:        make(map[time.Time]bool),
			Expires:    exp,
			LastUpdate: time.Now().In(location),
		}
		log.Println("init counter state")

		return
	}

	counter = *val
	log.Printf("restore counter state: %v\n", counter)
}

// GetState retrieve counter state
func GetState() (models.State, error) {
	defer recoverPanic()

	mutex.Lock()

	now := time.Now().In(location)
	state := counter.Inc(now)

	err := dal.Store(&counter, store)
	if err != nil {
		log.Println("can't write to strore", err)
	}

	mutex.Unlock()

	return state, nil
}

func recoverPanic() {
	if r := recover(); r != nil {
		log.Println("recovered", r)
	}
}
