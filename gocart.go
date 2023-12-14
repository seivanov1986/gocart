package gocart

import (
	"net/http"

	"github.com/seivanov1986/gocart/internal/http/user"
)

func UserCreateHttpHandler() func(http.ResponseWriter, *http.Request) {
	return user.New().Create
}
