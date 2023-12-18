package example

import (
	"context"

	"github.com/seivanov1986/gocart/pkg/cache"
)

const (
	resultString = "example widget"
)

func Widget(ctx context.Context, resources cache.BuilderResources) (*string, error) {
	result := resultString

	return &result, nil
}
