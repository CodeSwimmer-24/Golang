package middleware

import (
	"fmt"
	"net/http"
)

// Example middleware to log requests
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Received %s request for %s\n", r.Method, r.URL)
		next.ServeHTTP(w, r)
	})
}
