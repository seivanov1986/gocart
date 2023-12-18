package manager

import (
	"github.com/seivanov1986/gocart"
	"github.com/seivanov1986/gocart/pkg/cache"
)

type widgetManager struct {
	widgets   map[string]gocart.Widget
	resources cache.BuilderResources
}

func New() *widgetManager {
	return &widgetManager{
		widgets: map[string]gocart.Widget{},
	}
}
