package auth

import (
	"time"
)

const (
	NeverExpire    = -1 // NeverExpire 永不过期
	NotValueExpire = -2 // NotValueExpire 没有值过期
)

type Store interface {
	Get(key []byte) ([]byte, error)
	Put(key []byte, value []byte, timeout time.Duration) error

	Delete(key []byte) error
	Update(key []byte, value []byte) error

	CheckTimeout(key []byte) (time.Duration, error)
	UpdateTimeout(key []byte, timeout time.Duration) error
}
