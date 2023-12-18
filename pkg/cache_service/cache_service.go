package cache_service

import (
	"context"
	"github.com/seivanov1986/gocart"
)

type CacheService interface {
	Make()
}

type cacheService struct {
	cacheBuilder gocart.CacheBuilder
}

func New(cacheBuilder gocart.CacheBuilder) *cacheService {
	return &cacheService{cacheBuilder: cacheBuilder}
}

func (c *cacheService) Make() {
	c.cacheBuilder.Pages(context.Background())
}
