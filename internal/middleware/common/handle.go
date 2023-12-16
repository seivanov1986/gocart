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

		ctx := context.WithValue(r.Context(), "service_base_path", a.ServiceBasePath)
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
