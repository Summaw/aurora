package middleware

import (
	"time"

	"github.com/Summaw/aurora"
)

type FiberContext interface {
	Method() string
	Path() string
	IP() string
	Next() error
	Response() FiberResponse
}

type FiberResponse interface {
	StatusCode() int
}

func Fiber(log *aurora.Logger) func(FiberContext) error {
	return func(c FiberContext) error {
		start := time.Now()

		err := c.Next()

		status := c.Response().StatusCode()
		latency := time.Since(start)

		entry := log.Info("HTTP Request")
		if status >= 400 {
			entry = log.Warn("HTTP Request")
		}
		if status >= 500 {
			entry = log.Error("HTTP Request")
		}

		entry.
			Str("method", c.Method()).
			Str("path", c.Path()).
			Int("status", status).
			Dur("latency", latency).
			Str("ip", c.IP()).
			Send()

		return err
	}
}

type HTTPHandler struct {
	Logger *aurora.Logger
}

func NewHTTPHandler(log *aurora.Logger) *HTTPHandler {
	return &HTTPHandler{Logger: log}
}

func (h *HTTPHandler) LogRequest(method, path, ip string, status int, latency time.Duration) {
	entry := h.Logger.Info("HTTP Request")
	if status >= 400 {
		entry = h.Logger.Warn("HTTP Request")
	}
	if status >= 500 {
		entry = h.Logger.Error("HTTP Request")
	}

	entry.
		Str("method", method).
		Str("path", path).
		Int("status", status).
		Dur("latency", latency).
		Str("ip", ip).
		Send()
}
