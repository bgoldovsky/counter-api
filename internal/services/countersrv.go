package services

import (
	"errors"
	"log"
	"time"

	"github.com/bgoldovsky/counter-api/internal/dal"
	"github.com/bgoldovsky/counter-api/internal/models"
)

// CounterService asynchronous counter service
type CounterService struct {
	counter models.Requests
	store   string
	loc     *time.Location

	errChan   chan error
	incChan   chan struct{}
	stateChan chan models.State
}

// NewCounter CounterService constructor
func NewCounter(filepath string, expires int) (Counter, error) {
	if filepath == "" {
		err := errors.New("store path not specified")
		return nil, err
	}

	var loc *time.Location
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}

	val, err := dal.Get(filepath)
	if err != nil {
		val = &models.Requests{
			Map:        make(map[time.Time]bool),
			Expires:    expires,
			LastUpdate: time.Now().In(loc),
		}
	}

	c := CounterService{
		counter:   *val,
		store:     filepath,
		loc:       loc,
		errChan:   make(chan error, 1),
		incChan:   make(chan struct{}),
		stateChan: make(chan models.State),
	}

	go c.run()

	return &c, nil
}

// GetState retrieve counter state
func (c *CounterService) GetState() models.State {
	return <-c.stateChan
}

// Increment counter state
func (c *CounterService) Increment() error {
	c.incChan <- struct{}{}
	return <-c.errChan
}

func (c *CounterService) incrementCounter() error {
	now := time.Now().In(c.loc)
	c.counter.Increment(now)

	err := dal.Store(&c.counter, c.store)
	if err != nil {
		log.Println("can't write to strore", err)
		return err
	}

	return nil
}

func (c *CounterService) getState() models.State {
	now := time.Now().In(c.loc)
	return c.counter.State(now)
}

func (c *CounterService) run() {
	for {
		select {
		case <-c.incChan:
			c.errChan <- c.incrementCounter()
		case c.stateChan <- c.getState():
		}
	}
}
