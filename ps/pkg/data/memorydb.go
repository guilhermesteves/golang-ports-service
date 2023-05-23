package data

import (
	"fmt"
	"sync"
)

type (
	IDatabase interface {
		Create(key string, value interface{}) error
		Update(key string, value interface{}) error
		Select(key string) (interface{}, error)
		Delete(key string) error
	}
)

type db struct {
	memory sync.Map
}

func New() IDatabase {
	return &db{
		memory: sync.Map{},
	}
}

func (s *db) Create(key string, value interface{}) error {
	return s.write(key, value)
}

func (s *db) Update(key string, value interface{}) error {
	return s.write(key, value)
}

func (s *db) Select(key string) (interface{}, error) {
	return s.read(key)
}

func (s *db) Delete(key string) error {
	return nil
}

func (s *db) write(key string, value interface{}) error {
	if s == nil {
		return fmt.Errorf("database is null")
	}

	// delete for a possible update
	s.memory.LoadAndDelete(key)
	s.memory.Store(key, value)

	return nil
}

func (s *db) read(key string) (interface{}, error) {
	if s == nil {
		return nil, fmt.Errorf("database is null")
	}
	val, ok := s.memory.Load(key)
	if !ok {
		return nil, fmt.Errorf("not found (key: %s)", key)
	}
	return val, nil
}
