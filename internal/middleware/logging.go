package middleware

import (
	"project/internal/logging"
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {

		wrapped := &responseWriter{ResponseWriter: w, statusCode: 200}

		next.ServeHTTP(wrapped, req)

		logging.Log.Info(
			getRealIP(req),
			req.Method,
			wrapped.statusCode,
			req.URL.Path,
			req.Referer(),
		)
	})
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func getRealIP(req *http.Request) string {
	if ip := req.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	if ip := req.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	return req.RemoteAddr
}
