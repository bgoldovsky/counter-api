package dal

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"

	models "github.com/bgoldovsky/counter-api/internal/models"
)

// RepoGob repo gob implementation
type RepoGob struct {
	path string
}

// New creates new repo
func New(path string) (Repo, error) {
	if path == "" {
		return nil, errors.New("store path not specified")
	}

	dir, _ := filepath.Split(path)
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		return &RepoGob{path: path}, nil
	}

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return &RepoGob{path: path}, nil
}

// Save counter state to gob
func (r *RepoGob) Save(s *models.Counter) error {

	if s == nil {
		return errors.New("counter struct not specified")
	}

	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(s)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.path, buf.Bytes(), 0600)
	if err != nil {
		return err
	}

	return nil
}

// Get counter state from gob
func (r *RepoGob) Get() (*models.Counter, error) {
	raw, err := ioutil.ReadFile(r.path)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buf)

	c := new(models.Counter)
	err = dec.Decode(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
