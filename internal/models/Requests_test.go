package models_test

import (
	"testing"
	"time"

	. "github.com/bgoldovsky/counter-api/internal/models"
)

func TestIncSuccess(t *testing.T) {
	t.Parallel()

	r := Requests{}
	r.Map = make(map[time.Time]bool)
	r.Expires = 5

	const expected = 10

	for i := 1; i < 10; i++ {
		r.Increment(time.Now())
	}

	now := time.Now()

	r.Increment(now)
	act := r.State(now)

	if act.Total != expected {
		t.Errorf("counter not equals expected.\nact: %v\nexp %v\n", act.Total, expected)
	}
}

func TestIncExpired(t *testing.T) {
	t.Parallel()

	const expected = 1
	const expires = 1

	r := Requests{Expires: expires, Map: make(map[time.Time]bool)}

	for i := 0; i < 10; i++ {
		r.Increment(time.Now())
	}

	time.Sleep(time.Second * expires)
	now := time.Now()

	r.Increment(now)
	act := r.State(now)

	if act.Total != expected {
		t.Errorf("counter not equals expected.\nact: %v\nexp: %v\n", act.Total, expected)
	}
}

func TestStateSuccess(t *testing.T) {
	t.Parallel()

	r := Requests{}
	r.Map = make(map[time.Time]bool)
	r.Expires = 5

	const expected = 10

	for i := 0; i < 10; i++ {
		r.Increment(time.Now())
	}

	now := time.Now()

	act := r.State(now)

	if act.Total != expected {
		t.Errorf("counter not equals expected.\nact: %v\nexp %v\n", act.Total, expected)
	}
}

func TestStateExpired(t *testing.T) {
	t.Parallel()

	const expected = 0
	const expires = 1

	r := Requests{Expires: expires, Map: make(map[time.Time]bool)}

	for i := 0; i < 10; i++ {
		r.Increment(time.Now())
	}

	time.Sleep(time.Second * expires)
	now := time.Now()

	act := r.State(now)

	if act.Total != expected {
		t.Errorf("counter not equals expected.\nact: %v\nexp: %v\n", act.Total, expected)
	}
}
