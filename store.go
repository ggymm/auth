package auth

import (
	"time"

	"github.com/pkg/errors"
)

const (
	NeverExpire    = -1 // NeverExpire 永不过期
	NotValueExpire = -2 // NotValueExpire 没有值过期
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type Store interface {
	Get(key string) ([]byte, error)
	Put(key string, value []byte, timeout time.Duration) error

	Delete(key string) error
	Update(key string, value []byte) error

	CheckTimeout(key string) (time.Duration, error)
	UpdateTimeout(key string, timeout time.Duration) error
}
