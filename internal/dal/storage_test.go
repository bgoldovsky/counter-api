package dal

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/bgoldovsky/counter-api/internal/models"
)

func TestStorage_Success(t *testing.T) {
	const filepath = "./testfile.gob"

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		t.Error("can't load location")
	}

	defer clear(filepath)

	exp := models.Requests{Map: map[time.Time]bool{
		time.Now().In(loc): true,
		time.Now().In(loc): true,
		time.Now().In(loc): true,
	}}

	err = Store(&exp, filepath)
	if err != nil {
		t.Error("can't store data from file")
	}

	act, err := Get(filepath)
	if err != nil {
		t.Error("can't store data from file")
	}

	if eq := reflect.DeepEqual(exp.Map, &act.Map); eq {
		println(exp.Map, act.Map)
		t.Errorf("state not equals. \nexpected: %+v. \nactial: %+v.", exp, act)
	}
}

func TestStorage_WrongTime(t *testing.T) {
	const filepath = "./testfile.gob"
	defer clear(filepath)

	now := time.Now()
	exp := models.Requests{Map: map[time.Time]bool{
		now: true,
	}}

	Store(&exp, filepath)
	act, _ := Get(filepath)

	if exp.Map[now] == act.Map[now] {
		t.Errorf("time is equals")
	}
}

func TestStorage_NoPath(t *testing.T) {
	const filepath = ""

	exp := models.Requests{Map: map[time.Time]bool{
		time.Now(): true,
	}}

	err := Store(&exp, filepath)
	if err == nil {
		t.Error("not detected: store path not specified")
	}
}

func TestStorage_NoObj(t *testing.T) {
	const filepath = "./testfile.gob"

	err := Store(nil, filepath)
	if err == nil {
		t.Error("not detected: store struct not specified")
	}
}

func clear(filepath string) {
	var err = os.Remove(filepath)
	if err != nil {
		panic(err)
	}
}
