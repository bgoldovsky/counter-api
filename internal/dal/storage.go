package dal

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"

	models "github.com/bgoldovsky/counter-api/internal/models"
)

// Store counter state to gob
func Store(r *models.Requests, filename string) (err error) {

	if filename == "" {
		return errors.New("store path not specified")
	}

	if r == nil {
		return errors.New("store struct not specified")
	}

	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)

	err = enc.Encode(r)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, buf.Bytes(), 0600)
	if err != nil {
		return err
	}

	return
}

// Get counter state from gob
func Get(filename string) (r *models.Requests, err error) {
	if filename == "" {
		return nil, errors.New("store path not specified")
	}

	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buf)

	err = dec.Decode(&r)
	if err != nil {
		return nil, err
	}

	return
}
