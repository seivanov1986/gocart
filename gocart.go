package gocart

import (
	"github.com/seivanov1986/sql_client"

	"github.com/seivanov1986/gocart/internal/http/auth"
	"github.com/seivanov1986/gocart/internal/http/file"
	"github.com/seivanov1986/gocart/internal/http/user"
	auth2 "github.com/seivanov1986/gocart/internal/middleware/auth"
	"github.com/seivanov1986/gocart/internal/middleware/common"
	"github.com/seivanov1986/gocart/internal/middleware/cors"
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
		database:       options.database,
		sessionManager: options.sessionManager,
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

func (g *goCart) FileHandler() file.Handle {
	return file.New()
}

func (g *goCart) AuthMiddleware() auth2.Middleware {
	g.checkSessionManager()
	return auth2.New(g.sessionManager)
}

func (g *goCart) CommonMiddleware(serviceBasePath string) common.Middleware {
	return common.New(serviceBasePath)
}

func (g *goCart) CorsMiddleware() cors.Middleware {
	return cors.New()
}

func (g *goCart) checkDatabase() {
	if g.database == nil {
		panic("database must be an object")
	}
}

func (g *goCart) checkSessionManager() {
	if g.sessionManager == nil {
		panic("session manager must be an object")
	}
}
