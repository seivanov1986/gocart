package gocart

import (
	"github.com/seivanov1986/gocart/internal/http/user"
)

func UserHttpHandler() user.User {
	return user.New()
}
