package cache

import (
	"context"

	"github.com/seivanov1986/gocart"
	"github.com/seivanov1986/gocart/internal/repository"
	"github.com/seivanov1986/gocart/internal/repository/sefurl"
)

type builder struct {
	hub    repository.Hub
	widget gocart.Widget
}

func NewBuilder(hub repository.Hub, widget gocart.Widget) *builder {
	return &builder{hub: hub, widget: widget}
}

func (b *builder) Pages(ctx context.Context) ([]sefurl.SefUrlListRow, error) {
	return nil, nil
}

func (b *builder) Handler(ctx context.Context, pages []sefurl.SefUrlListRow) error {
	return nil
}
