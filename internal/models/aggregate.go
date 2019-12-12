package models

import (
	"time"
)

// Counter is a map and meta of request times
// Not thread safe
type Counter struct {
	Map        map[time.Time]bool
	Expires    int
	LastUpdate time.Time
}

// NewCounter creates new request counter
func NewCounter(expires int, lastUpdate time.Time) *Counter {
	return &Counter{
		Map:        make(map[time.Time]bool),
		Expires:    expires,
		LastUpdate: lastUpdate,
	}
}

// Inc request counter
func (c *Counter) Inc(now time.Time) {
	c.Map[now] = true
	c.LastUpdate = now
}

// State returns state of request counter
func (c *Counter) State(now time.Time) State {
	c.clear(now)

	return State{Total: len(c.Map), LastUpdate: c.LastUpdate}
}

func (c *Counter) clear(now time.Time) {
	for key := range c.Map {
		if now.Sub(key).Seconds() > float64(c.Expires) {
			delete(c.Map, key)
		}
	}
}
