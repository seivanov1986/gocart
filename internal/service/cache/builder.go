package cache

import (
	"context"

	"github.com/seivanov1986/gocart"
	"github.com/seivanov1986/gocart/internal/repository"
	"github.com/seivanov1986/gocart/internal/repository/sefurl"
)

type builder struct {
	hub           repository.Hub
	widgetManager gocart.WidgetManager
}

func NewBuilder(hub repository.Hub, widgetManager gocart.WidgetManager) *builder {
	return &builder{hub: hub, widgetManager: widgetManager}
}

func (b *builder) Pages(ctx context.Context) ([]sefurl.SefUrlListRow, error) {
	return nil, nil
}

func (b *builder) Handler(ctx context.Context, pages []sefurl.SefUrlListRow) error {
	return nil
}
