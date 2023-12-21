package observer

import (
	"context"
)

const (
	serviceBasePath = "service_base_path"
)

type observer struct {
}

func New() *observer {
	return &observer{}
}

func (o *observer) GetServiceBasePath(ctx context.Context) string {
	result := ""

	if value, ok := ctx.Value(serviceBasePath).(string); ok {
		result = value
	}

	return result
}

func (o *observer) SetServiceBasePath(ctx context.Context, path string) context.Context {
	return context.WithValue(ctx, serviceBasePath, path)
}
