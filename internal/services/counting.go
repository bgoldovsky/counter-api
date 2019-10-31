package services

import (
	"errors"
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

// Init load current state from storage
func Init(path string, exp int) error {
	if path == "" {
		return errors.New("store path not specified")
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

		return nil
	}

	counter = *val

	log.Printf("restore counter state: %v\n", counter)

	return nil
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
