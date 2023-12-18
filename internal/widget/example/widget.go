package example

import (
	"context"
)

const (
	resultString = "example widget"
)

func Widget(ctx context.Context) (*string, error) {
	result := resultString

	return &result, nil
}
