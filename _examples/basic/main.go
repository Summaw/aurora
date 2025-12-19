package main

import (
	"time"

	"github.com/Summaw/aurora"
)

func main() {
	aurora.Banner("MYAPP").
		Gradient("cyberpunk").
		Tagline("My Awesome Application").
		Version("v1.0.0").
		Render()

	log := aurora.New()

	log.Info("Application started").Send()
	log.Success("Database connected").Str("host", "localhost:5432").Send()
	log.Warn("Cache miss rate high").Int("percent", 23).Send()
	log.Error("Connection timeout").Str("service", "redis").Send()

	log.Info("Request completed").
		Str("method", "POST").
		Str("path", "/api/users").
		Int("status", 201).
		Dur("latency", 23*time.Millisecond).
		Send()
}
