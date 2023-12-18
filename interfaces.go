package gocart

import (
	"context"
	"time"
)

type SessionManager interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Exists(keys ...string) (bool, error)
	Del(keys ...string) (bool, error)
}

type UrlListRow struct {
	ID       int64
	Url      string
	Path     string
	Name     string
	Type     int64
	IdObject int64
}

type CacheBuilder interface {
	Pages(ctx context.Context) ([]UrlListRow, error)
	Handler(ctx context.Context, pages []UrlListRow) error
}

type WidgetManager interface {
	Render(ctx context.Context, name string) (*string, error)
	Register(name string, widget Widget)
}

type Widget interface {
	Execute() (*string, error)
}
