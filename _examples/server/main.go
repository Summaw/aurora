package main

import (
	"time"

	"github.com/Summaw/aurora"
)

func main() {
	aurora.Banner("API SERVER").
		Gradient("ocean").
		Tagline("RESTful API Service").
		Version("v2.0.0").
		Render()

	aurora.Divider("Configuration").Gradient("ocean").Render()

	aurora.KV(map[string]any{
		"Environment": "production",
		"Go Version":  "1.21.0",
		"Port":        8080,
		"TLS":         true,
	}).Gradient("sunset").Render()

	aurora.Divider("Starting Services").Gradient("ocean").Render()

	log := aurora.New(aurora.WithCaller(true))

	log.Info("Loading configuration").Str("file", "config.yaml").Send()
	log.Success("Database pool initialized").Int("connections", 10).Send()
	log.Success("Redis cache connected").Str("host", "localhost:6379").Send()
	log.Info("Starting HTTP server").Str("addr", ":8080").Send()

	aurora.Divider("Service Status").Gradient("ocean").Render()

	aurora.Table(
		[]string{"Service", "Status", "Latency"},
		[][]string{
			{"api-gateway", "● Healthy", "12ms"},
			{"postgres", "● Healthy", "3ms"},
			{"redis", "● Healthy", "1ms"},
			{"elasticsearch", "⚠ Degraded", "89ms"},
		},
	).Gradient("mint").Render()

	aurora.Divider("Incoming Requests").Gradient("ocean").Render()

	log.Info("HTTP Request").
		Str("method", "GET").
		Str("path", "/api/v1/users").
		Int("status", 200).
		Dur("latency", 15*time.Millisecond).
		Send()

	log.Info("HTTP Request").
		Str("method", "POST").
		Str("path", "/api/v1/orders").
		Int("status", 201).
		Dur("latency", 45*time.Millisecond).
		Send()

	log.Warn("HTTP Request").
		Str("method", "GET").
		Str("path", "/api/v1/products/999").
		Int("status", 404).
		Dur("latency", 8*time.Millisecond).
		Send()

	log.Error("HTTP Request").
		Str("method", "POST").
		Str("path", "/api/v1/payments").
		Int("status", 500).
		Str("error", "payment gateway timeout").
		Dur("latency", 30000*time.Millisecond).
		Send()

	aurora.Divider("").Gradient("ocean").Render()
}
