package gocart

import (
	"context"
	"time"

	"github.com/seivanov1986/gocart/internal/repository/sefurl"
)

type SessionManager interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Exists(keys ...string) (bool, error)
	Del(keys ...string) (bool, error)
}

type CacheBuilder interface {
	Pages(ctx context.Context) ([]sefurl.SefUrlListRow, error)
	Handler(ctx context.Context, pages []sefurl.SefUrlListRow) error
}

type Widget interface {
	Render(ctx context.Context, name string) (*string, error)
}
