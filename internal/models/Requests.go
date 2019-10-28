package models

import "time"

// Requests is a map of request times
type Requests struct {
	Map map[time.Time]bool
}

// Inc increment state of counter
func (r *Requests) Inc(now time.Time, expires int) {
	r.Map[now] = true
	for key := range r.Map {
		if now.Sub(key).Seconds() > float64(expires) {
			delete(r.Map, key)
		}
	}
}

// Count returns amount of counter
func (r *Requests) Count() int {
	return len(r.Map)
}
