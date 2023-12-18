package cache

import (
	"context"

	"github.com/seivanov1986/gocart"
	"github.com/seivanov1986/gocart/internal/repository"
	"github.com/seivanov1986/gocart/internal/repository/sefurl"
)

type BuilderResources struct {
	hub       repository.Hub
	schemaOrg SchemaOrg
	assets    Assets
}

type builder struct {
	resources     BuilderResources
	widgetManager gocart.WidgetManager
}

func NewBuilder(hub repository.Hub, widgetManager gocart.WidgetManager) *builder {
	resources := BuilderResources{
		hub:       hub,
		schemaOrg: NewSchemaOrg(),
		assets:    NewAsset(),
	}

	widgetManager.SetResources(resources)
	return &builder{resources: resources, widgetManager: widgetManager}
}

func (b *builder) Pages(ctx context.Context) ([]sefurl.SefUrlListRow, error) {
	return nil, nil
}

func (b *builder) Handler(ctx context.Context, pages []sefurl.SefUrlListRow) error {
	b.widgetManager.Render(ctx, "example")

	return nil
}
