package cache_builder

import (
	"context"

	"github.com/seivanov1986/gocart/client"
)

type Service interface {
	Make(ctx context.Context) error
	MakeObject(ctx context.Context, row client.SefUrlItem) ([]byte, error)
}
