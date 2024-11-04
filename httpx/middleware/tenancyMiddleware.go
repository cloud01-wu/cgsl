package middleware

import (
	"net/http"

	"github.com/gorilla/context"
)

func TenancyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// tenancy identity forwarded by API gateway
		tenancyString := r.Header.Get("X-User-Tenancy")

		context.Set(r, "tenancy", tenancyString)

		next.ServeHTTP(w, r)
	})
}
