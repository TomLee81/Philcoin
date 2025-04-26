package middleware

import (
    "net/http"
    "time"

    "golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 5)

// RateLimit applies simple token bucket per process
type RateLimit(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}