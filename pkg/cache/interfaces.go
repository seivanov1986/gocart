package cache

import (
	"context"
)

type Service interface {
	Make(ctx context.Context) error
}
