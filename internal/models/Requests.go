package models

import "time"

// Requests is a map of request times
type Requests struct {
	Map        map[time.Time]bool
	Expires    int
	LastUpdate time.Time
}

// Inc increment state of counter and returns state
func (r *Requests) Inc(now time.Time) State {
	r.Map[now] = true
	r.LastUpdate = now

	r.recalc(now)

	return State{Total: len(r.Map), LastUpdate: r.LastUpdate}
}

// State returns state without increment
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
