package services_test

import (
	"errors"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/bgoldovsky/counter-api/internal/dal/mock"
	"github.com/bgoldovsky/counter-api/internal/models"
	. "github.com/bgoldovsky/counter-api/internal/services"
	"github.com/golang/mock/gomock"
)

func TestServiceCreateSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockRepo(ctrl)

	repo.EXPECT().Get().Return(nil, errors.New("can't load counter"))

	_, err := NewCounterService(repo, 100)
	if err != nil {
		t.Error("creation service error:", err)
	}
}

func TestServiceLoadSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockRepo(ctrl)

	counter := models.NewCounter(100, time.Now())
	repo.EXPECT().Get().Return(counter, nil)

	_, err := NewCounterService(repo, 100)
	if err != nil {
		t.Error("creation service error:", err)
	}
}

func TestServiceCreateErrorRepo(t *testing.T) {
	exp := errors.New("repo not specified")

	_, err := NewCounterService(nil, 100)
	if !reflect.DeepEqual(exp, err) {
		t.Errorf("error expected %v, got %v", exp, err)
	}
}

func TestServiceCreateErrorExpires(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockRepo(ctrl)

	exp := errors.New("expires must be positive")
	data := []int{-1, 0}

	for _, expires := range data {
		_, err := NewCounterService(repo, expires)
		if !reflect.DeepEqual(exp, err) {
			t.Errorf("error expected %v, got %v", exp, err)
		}
	}
}

func TestIncSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockRepo(ctrl)
	counter := models.NewCounter(100, time.Now())

	repo.EXPECT().Get().Return(counter, nil)
	repo.EXPECT().Save(gomock.Any()).Return(nil)

	service, _ := NewCounterService(repo, 100)

	err := service.Inc()
	if err != nil {
		t.Error("increment error:", err)
	}
}

func TestIncError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockRepo(ctrl)
	counter := models.NewCounter(100, time.Now())

	exp := errors.New("save error")

	repo.EXPECT().Get().Return(counter, nil)
	repo.EXPECT().Save(gomock.Any()).Return(exp)

	service, _ := NewCounterService(repo, 100)

	err := service.Inc()
	if !reflect.DeepEqual(exp, err) {
		t.Errorf("error expected %v, got %v", exp, err)
	}
}

func TestStateSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockRepo(ctrl)
	counter := models.NewCounter(100, time.Now())

	exp := &models.State{Total: 1}

	repo.EXPECT().Get().Return(counter, nil)
	repo.EXPECT().Save(gomock.Any()).Return(nil)

	service, _ := NewCounterService(repo, 100)

	err := service.Inc()
	if err != nil {
		t.Fatal("increment error:", err)
	}

	act := service.State()
	if act.Total != exp.Total {
		t.Errorf("expected %v, got %v", exp, err)
	}
}

func TestRaceSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	repo := mock.NewMockRepo(ctrl)
	counter := models.NewCounter(100, time.Now())

	repo.EXPECT().Get().Return(counter, nil)

	service, _ := NewCounterService(repo, 100)

	var wg sync.WaitGroup
	wg.Add(2)

	for i := 0; i < 1000; i++ {
		repo.EXPECT().Save(gomock.Any()).Return(nil)
	}

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			service.Inc()
		}
	}()

	go func() {
		defer wg.Done()

		for i := 0; i < 1000; i++ {
			service.State()
		}
	}()

	wg.Wait()
}
