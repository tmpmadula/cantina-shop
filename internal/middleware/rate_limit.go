// internal/middleware/rate_limit.go
package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

func RateLimitMiddleware(next http.Handler) http.Handler {
	limiter := rate.NewLimiter(1, 5) // 1 request per second, burst of 5

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
