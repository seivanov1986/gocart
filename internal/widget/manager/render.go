package manager

import (
	"context"
	"fmt"
)

func (w *widgetManager) Render(ctx context.Context, name string) (*string, error) {
	widget, ok := w.widgets[name]
	if !ok {
		return nil, fmt.Errorf("widget not found")
	}

	return widget(ctx)
}
