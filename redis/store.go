package redis

import (
	"time"
)

type Store struct {
}

type Config struct {
	Dir string
}

func NewStore(config Config) (store *Store, err error) {
	store = &Store{}
	return
}

func (store *Store) Get(key []byte) ([]byte, error) {
	return nil, nil
}

func (store *Store) Put(key []byte, value []byte, timeout time.Duration) error {
	return nil
}

func (store *Store) Delete(key []byte) error {
	return nil
}

func (store *Store) Update(key []byte, value []byte) error {
	return nil
}

func (store *Store) CheckTimeout(key []byte) (time.Duration, error) {
	return 0, nil
}

func (store *Store) UpdateTimeout(key []byte, timeout time.Duration) error {
	return nil
}
