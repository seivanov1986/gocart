package manager

import (
	"context"
	"fmt"

	"github.com/seivanov1986/gocart/internal/service/cache"
)

func (w *widgetManager) Render(ctx context.Context, name string) (*string, error) {
	widget, ok := w.widgets[name]
	if !ok {
		return nil, fmt.Errorf("widget not found")
	}

	return widget(ctx, w.resources)
}

func (b *widgetManager) SetResources(resources cache.BuilderResources) {
	b.resources = resources
}
