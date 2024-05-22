// internal/middleware/ratelimit.go
package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimit(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(1, 3) // 1 request per second, with a burst size of 3

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
