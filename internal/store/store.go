package store

import "errors"

var ErrNotFound = errors.New("key not found")

type Store struct {
	data   map[string]string
	txStack []map[string]string
}

func New() *Store {
	return &Store{
		data: make(map[string]string),
	}
}

func (s *Store) Write(k, v string) {
	s.data[k] = v
}

func (s *Store) Read(k string) (string, error) {
	v, ok := s.data[k]
	if !ok {
		return "", ErrNotFound
	}
	return v, nil
}

func (s *Store) Delete(k string) error {
	if _, ok := s.data[k]; !ok {
		return ErrNotFound
	}
	delete(s.data, k)
	return nil
}

func (s *Store) Start() {
	snapshot := make(map[string]string, len(s.data))
	for k, v := range s.data {
		snapshot[k] = v
	}
	s.txStack = append(s.txStack, snapshot)
}

func (s *Store) Abort() error {
	if len(s.txStack) == 0 {
		return errors.New("no active transaction")
	}

	last := s.txStack[len(s.txStack)-1]
	s.txStack = s.txStack[:len(s.txStack)-1]
	s.data = last
	return nil
}

func (s *Store) Commit() error {
	if len(s.txStack) == 0 {
		return errors.New("no active transaction")
	}
	s.txStack = nil
	return nil
}

