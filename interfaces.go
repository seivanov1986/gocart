package gocart

import (
	"time"
)

type SessionManager interface {
	reDial() bool
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Exists(keys ...string) (bool, error)
	Del(keys ...string) (bool, error)
}
