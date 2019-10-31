package services

import (
	"os"
	"sync"
	"time"

	"testing"
)

func TestCounting_Success(t *testing.T) {
	const store = "./testfile.gob"
	const expiresSec = 60

	Init(store, expiresSec)
	defer clear(store)

	state, err := GetState()
	if err != nil {
		t.Error("can't get state: " + err.Error())
	}

	if state.Total != 1 {
		t.Errorf("invalid state: %d\n", state.Total)
	}
}

func TestCounting_Concurrency_Success(t *testing.T) {
	const store = "./testfile.gob"
	const expiresSec = 60

	Init(store, expiresSec)
	defer clear(store)

	runConcurrency()

	s, _ := GetState()
	if s.Total != 1000 {
		t.Errorf("invalid state: %d\n", s.Total)
	}
}

func TestCounting_Expired_Success(t *testing.T) {
	const store = "./testfile.gob"
	const expiresSec = 3

	Init(store, expiresSec)
	defer clear(store)

	runConcurrency()
	time.Sleep(time.Second * expiresSec)

	s, _ := GetState()
	if s.Total != 1 {
		t.Errorf("invalid state: %d\n", s.Total)
	}
}

func TestCounting_Expired_Fail(t *testing.T) {
	const store = "./testfile.gob"
	const expiresSec = 10

	Init(store, expiresSec)
	defer clear(store)

	runConcurrency()
	time.Sleep(time.Second)

	s, _ := GetState()
	if s.Total == 1 {
		t.Errorf("invalid state: %d\n", s.Total)
	}
}

func runConcurrency() {
	var wg sync.WaitGroup
	wg.Add(999)

	for i := 0; i < 999; i++ {
		go func() {
			defer wg.Done()
			GetState()
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
