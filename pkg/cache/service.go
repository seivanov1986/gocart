package cache

import (
	"github.com/seivanov1986/gocart"
)

type service struct {
	cacheBuilder gocart.CacheBuilder
}

func New(cacheBuilder gocart.CacheBuilder) *service {
	return &service{cacheBuilder: cacheBuilder}
}
