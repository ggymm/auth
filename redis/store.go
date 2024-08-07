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

func (store *Store) Get(key string) ([]byte, error) {
	return nil, nil
}

func (store *Store) Put(key string, value []byte, timeout time.Duration) error {
	return nil
}

func (store *Store) Delete(key string) error {
	return nil
}

func (store *Store) Update(key string, value []byte) error {
	return nil
}

func (store *Store) CheckTimeout(key string) (time.Duration, error) {
	return 0, nil
}

func (store *Store) UpdateTimeout(key string, timeout time.Duration) error {
	return nil
}
