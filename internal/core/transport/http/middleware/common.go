package core_http_middleware

import (
	"net/http"
	"time"

	core_logger "github.com/Mihail4531/golang-todo/internal/core/logger"
	core_http_response "github.com/Mihail4531/golang-todo/internal/core/transport/http/response"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

const (
	requestIDHeader = "X-Request-ID"
)

func RequestID() Middleware {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(requestIDHeader)
			if requestId == "" {
				requestId = uuid.NewString()
			}
			r.Header.Set(requestIDHeader, requestId)
			w.Header().Set(requestIDHeader, requestId)
			next.ServeHTTP(w, r)
		})
	}
}
func Logger(log *core_logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestId := r.Header.Get(requestIDHeader)
			l := log.With(
				zap.String("requestiD", requestId),
				zap.String("url", r.URL.String()),
			)
			ctx := core_logger.ToContext(r.Context(), l)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
func Panic() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			responseHandler := core_http_response.NewHTTPResponseHandler(log, w)
			defer func() {
				if p := recover(); p != nil {
					responseHandler.PanicResponse(p, "during handle Http request got unxpected panic")
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := core_logger.FromContext(ctx)
			rw := core_http_response.NewResponseWriter(w)
			before := time.Now()
			log.Debug("incoming http request", zap.Time("time", before.UTC()), zap.String("Method", r.Method))
			next.ServeHTTP(rw, r)
			log.Debug(
				"done http request",
				zap.Int("staus_code", rw.GetStatusCode()),
				zap.Duration("latency", time.Now().Sub(before)),
			)
		})
	}
}
