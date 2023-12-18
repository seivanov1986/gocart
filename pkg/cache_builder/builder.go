package cache_builder

import (
	"context"

	"github.com/seivanov1986/gocart"
	"github.com/seivanov1986/gocart/internal/repository"
)

type builder struct {
	hub repository.Hub
	widgetManager gocart.WidgetManager
}

func NewBuilder(hub repository.Hub, widgetManager gocart.WidgetManager) *builder {
	return &builder{hub: hub, widgetManager: widgetManager}
}

func (b *builder) RegisterWidget(name string, widget gocart.Widget) {
	b.widgetManager.Register(name, widget)
}

func (b *builder) Pages(ctx context.Context) ([]gocart.UrlListRow, error) {
	return nil, nil
}

func (b *builder) Handler(ctx context.Context, pages []gocart.UrlListRow) error {
	return nil
}
