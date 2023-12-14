package gocart

import (
	"github.com/seivanov1986/sql_client"

	"github.com/seivanov1986/gocart/internal/http/user"
	"github.com/seivanov1986/gocart/internal/repository"
	user2 "github.com/seivanov1986/gocart/internal/service/user"
)

type Options struct {
	database sql_client.DataBase
}

type OptionFunc func(*Options)

func WithDatabase(database sql_client.DataBase) OptionFunc {
	return func(o *Options) {
		o.database = database
	}
}

type goCart struct {
	database sql_client.DataBase
}

func New(opts ...OptionFunc) *goCart {
	options := Options{}
	for _, opt := range opts {
		opt(&options)
	}

	return &goCart{
		database: options.database,
	}
}

func (g *goCart) UserHttpHandler() user.User {
	if g.database == nil {
		panic("database must be an object")
	}

	hub := repository.New(g.database)
	service := user2.New(hub)
	return user.New(service)
}
