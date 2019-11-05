package services_test

import (
	"os"
	"sync"
	"testing"
	"time"

	. "github.com/bgoldovsky/counter-api/internal/services"
)

func TestCountingSuccess(t *testing.T) {
	const store = "./testfile.gob"
	const expiresSec = 60

	srv, err := NewCounter(store, expiresSec)
	if err != nil {
		t.Fatal("can't create service")
	}

	defer clear(store)

	srv.Increment()
	state := srv.GetState()

	if state.Total != 1 {
		t.Errorf("invalid state: %d\n", state.Total)
	}
}

func TestCountingConcurrencySuccess(t *testing.T) {

	const store = "./testfile.gob"
	const expiresSec = 60

	srv, err := NewCounter(store, expiresSec)
	if err != nil {
		t.Fatal("can't create service")
	}

	defer clear(store)

	runConcurrency(srv)

	s := srv.GetState()

	if s.Total != 1000 {
		t.Errorf("invalid state: %d\n", s.Total)
	}
}

func TestCountingExpiredSuccess(t *testing.T) {
	const store = "./testfile.gob"
	const expiresSec = 3

	srv, err := NewCounter(store, expiresSec)
	if err != nil {
		t.Fatal("can't create service")
	}
	defer clear(store)

	runConcurrency(srv)
	time.Sleep(time.Second * expiresSec)

	srv.Increment()

	s := srv.GetState()
	if s.Total != 1 {
		t.Errorf("invalid state: %d\n", s.Total)
	}
}

func TestCountingExpiredFail(t *testing.T) {
	const store = "./testfile.gob"
	const expiresSec = 10

	srv, err := NewCounter(store, expiresSec)
	if err != nil {
		t.Fatal("can't create service")
	}
	defer clear(store)

	runConcurrency(srv)
	time.Sleep(time.Second)

	s := srv.GetState()
	if s.Total == 1 {
		t.Errorf("invalid state: %d\n", s.Total)
	}
}

func runConcurrency(c Counter) {
	var wg sync.WaitGroup
	wg.Add(1000)

	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			c.Increment()
		}()
	}

	wg.Wait()
}

func clear(store string) {
	var err = os.Remove(store)
	if err != nil {
		panic(err)
	}
}
