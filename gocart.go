package gocart

import (
	"github.com/seivanov1986/sql_client"

	"github.com/seivanov1986/gocart/internal/http/auth"
	"github.com/seivanov1986/gocart/internal/http/user"
	"github.com/seivanov1986/gocart/internal/repository"
	authService "github.com/seivanov1986/gocart/internal/service/auth"
	user2 "github.com/seivanov1986/gocart/internal/service/user"
)

type Options struct {
	database       sql_client.DataBase
	sessionManager SessionManager
}

type OptionFunc func(*Options)

func WithDatabase(database sql_client.DataBase) OptionFunc {
	return func(o *Options) {
		o.database = database
	}
}

func WithSessionManager(sessionManager SessionManager) OptionFunc {
	return func(o *Options) {
		o.sessionManager = sessionManager
	}
}

type goCart struct {
	database       sql_client.DataBase
	sessionManager SessionManager
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

func (g *goCart) UserHttpHandler() user.Handle {
	g.checkDatabase()

	hub := repository.New(g.database)
	service := user2.New(hub)
	return user.New(service)
}

func (g *goCart) AuthHandler() auth.Handle {
	g.checkDatabase()
	g.checkSessionManager()

	hub := repository.New(g.database)
	service := authService.New(hub, g.sessionManager)
	return auth.New(service)
}

func (g *goCart) checkDatabase() {
	if g.database == nil {
		panic("database must be an object")
	}
}

func (g *goCart) checkSessionManager() {
	if g.database == nil {
		panic("database must be an object")
	}
}
