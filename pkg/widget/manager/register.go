package manager

import (
	"github.com/seivanov1986/gocart"
)

func (w *widgetManager) Register(name string, widget gocart.Widget) {
	w.widgets[name] = widget
}
