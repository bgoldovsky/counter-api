package models_test

import (
	"testing"
	"time"

	. "github.com/bgoldovsky/counter-api/internal/models"
)

func TestIncSuccess(t *testing.T) {
	t.Parallel()

	c := NewCounter(100, time.Now().UTC())
	c.Map = make(map[time.Time]bool)
	c.Expires = 5

	const exp = 10

	for i := 1; i < 10; i++ {
		c.Inc(time.Now())
	}

	now := time.Now()

	c.Inc(now)
	act := c.State(now)

	if act.Total != exp {
		t.Errorf("counter not equals expected. exp: %v, got %v\n", exp, act.Total)
	}
}

func TestIncExpired(t *testing.T) {
	t.Parallel()

	const expected = 1
	const expires = 1

	c := Counter{
		Expires: expires,
		Map:     make(map[time.Time]bool),
	}

	for i := 0; i < 10; i++ {
		c.Inc(time.Now())
	}

	time.Sleep(time.Second * expires)
	now := time.Now()

	c.Inc(now)
	act := c.State(now)

	if act.Total != expected {
		t.Errorf("counter not equals expected. exp: %v, got %v\n", expected, act.Total)
	}
}

func TestStateSuccess(t *testing.T) {
	t.Parallel()

	c := Counter{
		Expires: 5,
		Map:     make(map[time.Time]bool),
	}

	const exp = 10

	for i := 0; i < 10; i++ {
		c.Inc(time.Now())
	}

	now := time.Now()

	act := c.State(now)

	if act.Total != exp {
		t.Errorf("counter not equals expected. exp: %v, got %v\n", exp, act.Total)
	}
}

func TestStateExpired(t *testing.T) {
	t.Parallel()

	const expected = 0
	const expires = 1

	r := Counter{Expires: expires, Map: make(map[time.Time]bool)}

	for i := 0; i < 10; i++ {
		r.Inc(time.Now())
	}

	time.Sleep(time.Second * expires)
	now := time.Now()

	act := r.State(now)

	if act.Total != expected {
		t.Errorf("counter not equals expected. exp: %v, got %v\n", expected, act.Total)
	}
}
