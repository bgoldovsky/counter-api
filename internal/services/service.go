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
	counter models.Counter
	repo    dal.Repo

	errChan   chan error
	incChan   chan struct{}
	stateChan chan models.State
}

// NewCounterService CounterService constructor
func NewCounterService(repo dal.Repo, expires int) (*CounterService, error) {
	if repo == nil {
		return nil, errors.New("repo not specified")
	}

	if expires <= 0 {
		return nil, errors.New("expires must be positive")
	}

	val, err := repo.Get()
	if err != nil {
		log.Printf("can't load counter: %v", err)
		val = models.NewCounter(expires, time.Now().UTC())
	}

	c := CounterService{
		counter: *val,
		repo:    repo,

		errChan:   make(chan error, 1),
		incChan:   make(chan struct{}),
		stateChan: make(chan models.State),
	}

	go c.run()

	return &c, nil
}

// State actualize and retrieve counter state
func (c *CounterService) State() models.State {
	return <-c.stateChan
}

// Inc counter state
func (c *CounterService) Inc() error {
	c.incChan <- struct{}{}
	return <-c.errChan
}

func (c *CounterService) inc() error {
	now := time.Now().UTC()
	c.counter.Inc(now)

	err := c.repo.Save(&c.counter)
	if err != nil {
		log.Println("can't save to repo", err)
		return err
	}

	return nil
}

func (c *CounterService) state() models.State {
	now := time.Now().UTC()
	return c.counter.State(now)
}

func (c *CounterService) run() {
	for {
		select {
		case <-c.incChan:
			c.errChan <- c.inc()
		case c.stateChan <- c.state():
		}
	}
}
