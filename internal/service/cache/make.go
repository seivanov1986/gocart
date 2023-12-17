package cache

import (
	"context"
)

func (s *service) Make(ctx context.Context) error {
	pages, err := s.cacheBuilder.Pages(ctx)
	if err != nil {
		return err
	}
	
	return s.cacheBuilder.Handler(ctx, pages)
}
