# Aurora

**Beautiful, enterprise-grade console logging for Go.**

[![Go Reference](https://pkg.go.dev/badge/github.com/Summaw/aurora.svg)](https://pkg.go.dev/github.com/Summaw/aurora)
[![Go Report Card](https://goreportcard.com/badge/github.com/Summaw/aurora)](https://goreportcard.com/report/github.com/Summaw/aurora)

Aurora is a zero-dependency logging library that makes your terminal output stunning. Features gradient-colored ASCII banners, structured logging with tree-style output, progress bars, spinners, tables, and more.

## Installation

```bash
go get github.com/Summaw/aurora
```

## Quick Start

```go
package main

import "github.com/Summaw/aurora"

func main() {
    aurora.Banner("MYAPP").
        Gradient("cyberpunk").
        Tagline("My Application").
        Version("v1.0.0").
        Render()

    log := aurora.New()

    log.Info("Application started").Send()
    log.Success("Database connected").Str("host", "localhost").Send()
    log.Warn("High memory usage").Int("percent", 85).Send()
    log.Error("Connection failed").Err(err).Send()
}
```

## Features

### ASCII Art Banners

```go
aurora.Banner("AURORA").
    Font("block").           // block, slant, minimal
    Gradient("cyberpunk").   // 20+ built-in gradients
    Tagline("My App").
    Version("v1.0.0").
    Border("rounded").       // rounded, sharp, double, heavy, ascii
    Render()
```

### Structured Logging

```go
log := aurora.New()

log.Info("Request completed").
    Str("method", "POST").
    Str("path", "/api/users").
    Int("status", 201).
    Dur("latency", time.Since(start)).
    Send()
```

Output:
```
  14:23:01.123  ● INFO  Request completed
                ├─ method: POST
                ├─ path: /api/users
                ├─ status: 201
                └─ latency: 23.45ms
```

### Log Levels

```go
log.Trace("Trace message")       // ◦ gray
log.Debug("Debug message")       // ● dark gray
log.Info("Info message")         // ● blue
log.Success("Success message")   // ✓ green
log.Warn("Warning message")      // ⚠ yellow
log.Error("Error message")       // ✖ red
log.Fatal("Fatal message")       // 💀 bright red (exits)
```

### Configuration

```go
log := aurora.New(
    aurora.WithLevel(aurora.DebugLevel),
    aurora.WithCaller(true),
    aurora.WithJSON(true),  // JSON output for production
)
```

### UI Components

#### Tables

```go
aurora.Table(
    []string{"Service", "Status", "Latency"},
    [][]string{
        {"api", "● Healthy", "12ms"},
        {"db", "● Healthy", "3ms"},
    },
).Gradient("mint").Render()
```

#### Progress Bars

```go
bar := aurora.Progress("Downloading", 100)
for i := 0; i <= 100; i++ {
    bar.Set(i)
    time.Sleep(50 * time.Millisecond)
}
bar.Done()
```

#### Spinners

```go
spin := aurora.Spin("Connecting...")
// ... work ...
spin.Success("Connected!")
```

#### Dividers & Key-Value Display

```go
aurora.Divider("Configuration").Render()
aurora.KV(map[string]any{
    "Environment": "production",
    "Version": "1.0.0",
}).Render()
```

### Built-in Gradients

`aurora`, `sunset`, `ocean`, `neon`, `cyberpunk`, `miami`, `fire`, `forest`, `galaxy`, `retro`, `mint`, `peach`, `lavender`, `gold`, `ice`, `blood`, `matrix`, `vaporwave`, `rainbow`, `terminal`, `rose`, `sky`

### Custom Gradients

```go
aurora.Banner("APP").GradientRGB(
    aurora.RGB(255, 0, 128),
    aurora.RGB(0, 255, 255),
).Render()

aurora.Banner("APP").GradientMulti(
    "#ff0000", "#ff7f00", "#ffff00", "#00ff00",
).Render()
```

## Project Structure

```
aurora/
├── aurora.go           # Main public API
├── logger.go           # Logger implementation
├── entry.go            # Entry/field chaining
├── level.go            # Log levels
├── config.go           # Configuration
├── options.go          # Functional options
├── pkg/
│   ├── banner/         # ASCII banner system
│   ├── color/          # Color & gradient engine
│   └── style/          # UI components
├── middleware/         # HTTP middleware
├── docs/               # Documentation
└── _examples/          # Example code
```

## License

MIT License - see [LICENSE](LICENSE) for details.
