package manager

import (
	"github.com/seivanov1986/gocart"
)

type widgetManager struct {
	widgets map[string]gocart.Widget
}

func New() *widgetManager {
	return &widgetManager{
		widgets: map[string]gocart.Widget{},
	}
}
