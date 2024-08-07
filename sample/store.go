package main

import (
	"sync"
	"time"

	"github.com/ggymm/auth"
)

type Store struct {
	sync.Mutex
	data sync.Map
}

type StoreValue struct {
	Value  []byte
	Expire time.Time
}

func NewStore() (store *Store, err error) {
	store = &Store{}
	return
}

func (store *Store) read(key string) (*StoreValue, error) {
	store.Lock()
	defer store.Unlock()

	val, ok := store.data.Load(key)
	if !ok {
		return nil, auth.ErrKeyNotFound
	}

	data := val.(*StoreValue)
	if data.Expire.Before(time.Now()) {
		store.data.Delete(key)
		return nil, auth.ErrKeyNotFound
	}
	return data, nil
}

func (store *Store) write(key string, value *StoreValue) {
	store.Lock()
	defer store.Unlock()

	store.data.Store(key, value)
}

func (store *Store) Get(key string) ([]byte, error) {
	val, err := store.read(key)
	if err != nil {
		return nil, err
	}
	return val.Value, nil
}

func (store *Store) Put(key string, value []byte, timeout time.Duration) error {
	store.data.Store(key, &StoreValue{
		Value:  value,
		Expire: time.Now().Add(timeout),
	})

	time.AfterFunc(timeout, func() {
		val, ok := store.data.Load(key)
		if ok && val.(*StoreValue).Expire.Before(time.Now()) {
			store.data.Delete(key)
		}
	})
	return nil
}

func (store *Store) Delete(key string) error {
	store.data.Delete(key)
	return nil
}

func (store *Store) Update(key string, value []byte) error {
	val, err := store.read(key)
	if err != nil {
		return err
	}
	val.Value = value
	store.write(key, val)
	return nil
}

func (store *Store) CheckTimeout(key string) (time.Duration, error) {
	val, err := store.read(key)
	if err != nil {
		return 0, err
	}
	return val.Expire.Sub(time.Now()), nil
}

func (store *Store) UpdateTimeout(key string, timeout time.Duration) error {
	val, err := store.read(key)
	if err != nil {
		return err
	}
	val.Expire = time.Now().Add(timeout)
	store.write(key, val)
	return nil
}
