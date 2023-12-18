package gocart

import (
	"context"
	"time"

	"github.com/seivanov1986/gocart/pkg/cache"
)

type SessionManager interface {
	Set(key string, value interface{}, expiration time.Duration) error
	Get(key string) (string, error)
	Exists(keys ...string) (bool, error)
	Del(keys ...string) (bool, error)
}

type UrlListRow struct {
	ID       int64  `db:"id" json:"id"`
	Url      string `db:"url" json:"url"`
	Path     string `db:"path" json:"path"`
	Name     string `db:"name" json:"name"`
	Type     int64  `db:"type" json:"type"`
	IdObject int64  `db:"id_object" json:"id_object"`
}

type CacheBuilder interface {
	Pages(ctx context.Context) ([]UrlListRow, error)
	Handler(ctx context.Context, pages []UrlListRow) error
}

type WidgetManager interface {
	Render(ctx context.Context, name string) (*string, error)
	Register(name string, widget Widget)
	SetResources(resources cache.BuilderResources)
}

type Widget func(ctx context.Context, resources cache.BuilderResources) (*string, error)
