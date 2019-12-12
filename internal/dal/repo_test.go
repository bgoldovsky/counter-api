package dal_test

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	. "github.com/bgoldovsky/counter-api/internal/dal"
	"github.com/bgoldovsky/counter-api/internal/models"
)

func TestRepoCreateSuccess(t *testing.T) {
	const path = "./testfile.gob"

	_, err := New(path)
	if err != nil {
		t.Error("can't create repo", err)
	}
}

func TestRepoCreateNewDir(t *testing.T) {
	const path = "./new/testfile.gob"

	_, err := New(path)
	if err != nil {
		t.Error("can't create repo", err)
	}

	rmdir(path)
}

func TestRepoCreateErrorNoPath(t *testing.T) {
	_, err := New("")

	exp := errors.New("store path not specified")

	if !reflect.DeepEqual(exp, err) {
		t.Errorf("error expected %v, got %v", exp, err)
	}
}

func TestOperationsSuccess(t *testing.T) {
	const path = "./testfile.gob"

	defer clear(path)

	exp := models.Counter{Map: map[time.Time]bool{
		time.Now().UTC(): true,
		time.Now().UTC(): true,
		time.Now().UTC(): true,
	}}

	repo, err := New(path)

	err = repo.Save(&exp)
	if err != nil {
		t.Error("can't save data to file")
	}

	act, err := repo.Get()
	if err != nil {
		t.Error("can't load data from file")
	}

	if eq := reflect.DeepEqual(exp.Map, &act.Map); eq {
		println(exp.Map, act.Map)
		t.Errorf("state not equals. exp: %+v, got: %+v.", exp, act)
	}
}

func TestOperationsWrongTime(t *testing.T) {
	const path = "./testfile.gob"
	defer clear(path)

	repo, _ := New(path)

	now := time.Now()
	exp := models.Counter{Map: map[time.Time]bool{
		now: true,
	}}

	repo.Save(&exp)
	act, _ := repo.Get()

	if exp.Map[now] == act.Map[now] {
		t.Errorf("time is equals. exp %v, got %v", exp.Map[now], act.Map[now])
	}
}

func TestSaveErrorNoCounter(t *testing.T) {
	const path = "./testfile.gob"

	repo, _ := New(path)

	exp := errors.New("counter struct not specified")

	err := repo.Save(nil)
	if !reflect.DeepEqual(exp, err) {
		t.Errorf("error expected %v, got %v", exp, err)
	}
}

func TestGetErrorNoCounter(t *testing.T) {
	const path = "./testfile.gob"

	repo, _ := New(path)

	exp := fmt.Errorf("open %v: no such file or directory", path)

	_, err := repo.Get()
	if !reflect.DeepEqual(exp, err) {
		t.Errorf("error expected %v, got %v", exp, err)
	}
}

func clear(filepath string) {
	var err = os.Remove(filepath)
	if err != nil {
		panic(err)
	}
}

func rmdir(path string) {
	dir, _ := filepath.Split(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return
	}

	err := os.RemoveAll(dir)
	if err != nil {
		panic(err)
	}
}
