package widget_manager

import (
	"context"
	"fmt"

	"github.com/seivanov1986/gocart"
)

type widgetManager struct {
	widgets   map[string]gocart.Widget
}

func New() *widgetManager {
	return &widgetManager{
		widgets: map[string]gocart.Widget{},
	}
}

func (w *widgetManager) Register(name string, widget gocart.Widget) {
	w.widgets[name] = widget
}

func (w *widgetManager) Render(ctx context.Context, name string) (*string, error) {
	widget, ok := w.widgets[name]
	if !ok {
		return nil, fmt.Errorf("widget not found")
	}

	return widget.Execute()
}
