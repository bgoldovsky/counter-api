package models

import "time"

// Requests is a map and meta of request times
type Requests struct {
	Map        map[time.Time]bool
	Expires    int
	LastUpdate time.Time
}

// Increment request counter
func (r *Requests) Increment(now time.Time) {
	r.Map[now] = true
	r.LastUpdate = now

	r.recalc(now)
}

// State returns request counter
func (r *Requests) State(now time.Time) State {
	r.recalc(now)

	return State{Total: len(r.Map), LastUpdate: r.LastUpdate}
}

func (r *Requests) recalc(now time.Time) {
	for key := range r.Map {
		if now.Sub(key).Seconds() > float64(r.Expires) {
			delete(r.Map, key)
		}
	}
}
