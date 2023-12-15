package auth

import (
	"context"
	"net/http"

	"github.com/seivanov1986/gocart/helpers"
)

func (a middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")

		if auth == "" {
			helpers.HttpResponse(w, http.StatusUnauthorized)
			return
		}

		if !a.sessionClient.Exists(auth) {
			helpers.HttpResponse(w, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "authorization", auth)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}
