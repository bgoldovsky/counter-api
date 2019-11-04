package models_test

import (
	"testing"
	"time"

	. "github.com/bgoldovsky/counter-api/internal/models"
)

func TestInc_Success(t *testing.T) {
	r := Requests{}
	r.Map = make(map[time.Time]bool)
	r.Expires = 5

	const expected = 10

	for i := 1; i < 10; i++ {
		r.Inc(time.Now())
	}

	act := r.Inc(time.Now())

	if act.Total != expected {
		t.Errorf("counter not equals expected.\nact: %v\nexp %v\n", act.Total, expected)
	}
}

func TestInc_Expires(t *testing.T) {
	const expected = 1
	const expires = 1

	r := Requests{Expires: expires, Map: make(map[time.Time]bool)}

	for i := 0; i < 10; i++ {
		r.Inc(time.Now())
	}

	time.Sleep(time.Second * expires)

	act := r.Inc(time.Now())

	if act.Total != expected {
		t.Errorf("counter not equals expected.\nact: %v\nexp: %v\n", act.Total, expected)
	}
}

func TestState_Success(t *testing.T) {
	r := Requests{}
	r.Map = make(map[time.Time]bool)
	r.Expires = 5

	const expected = 10

	for i := 0; i < 10; i++ {
		r.Inc(time.Now())
	}

	act := r.State(time.Now())

	if act.Total != expected {
		t.Errorf("counter not equals expected.\nact: %v\nexp %v\n", act.Total, expected)
	}
}

func TestState_Expires(t *testing.T) {
	const expected = 0
	const expires = 1

	r := Requests{Expires: expires, Map: make(map[time.Time]bool)}

	for i := 0; i < 10; i++ {
		r.Inc(time.Now())
	}

	time.Sleep(time.Second * expires)

	act := r.State(time.Now())

	if act.Total != expected {
		t.Errorf("counter not equals expected.\nact: %v\nexp: %v\n", act.Total, expected)
	}
}
