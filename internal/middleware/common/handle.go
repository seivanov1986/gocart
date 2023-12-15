package common

import (
	"context"
	"net/http"
)

func (a middleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
			}
		}()

		ctx, cancel := context.WithCancel(r.Context())
		defer cancel()

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
