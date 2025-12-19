package aurora

import "github.com/Summaw/aurora/pkg/color"

type Level int

const (
	TraceLevel Level = iota
	DebugLevel
	InfoLevel
	SuccessLevel
	WarnLevel
	ErrorLevel
	FatalLevel
	PanicLevel
	Disabled
)

type LevelConfig struct {
	Name  string
	Icon  string
	Color color.RGB
	Bold  bool
}

var DefaultLevelConfigs = map[Level]LevelConfig{
	TraceLevel: {
		Name:  "TRACE",
		Icon:  "◦",
		Color: color.RGB{R: 128, G: 128, B: 128},
		Bold:  false,
	},
	DebugLevel: {
		Name:  "DEBUG",
		Icon:  "●",
		Color: color.RGB{R: 169, G: 169, B: 169},
		Bold:  false,
	},
	InfoLevel: {
		Name:  "INFO",
		Icon:  "●",
		Color: color.RGB{R: 96, G: 165, B: 250},
		Bold:  false,
	},
	SuccessLevel: {
		Name:  "SUCCESS",
		Icon:  "✓",
		Color: color.RGB{R: 74, G: 222, B: 128},
		Bold:  true,
	},
	WarnLevel: {
		Name:  "WARN",
		Icon:  "⚠",
		Color: color.RGB{R: 251, G: 191, B: 36},
		Bold:  true,
	},
	ErrorLevel: {
		Name:  "ERROR",
		Icon:  "✖",
		Color: color.RGB{R: 248, G: 113, B: 113},
		Bold:  true,
	},
	FatalLevel: {
		Name:  "FATAL",
		Icon:  "💀",
		Color: color.RGB{R: 239, G: 68, B: 68},
		Bold:  true,
	},
	PanicLevel: {
		Name:  "PANIC",
		Icon:  "🔥",
		Color: color.RGB{R: 185, G: 28, B: 28},
		Bold:  true,
	},
}

func (l Level) String() string {
	if cfg, ok := DefaultLevelConfigs[l]; ok {
		return cfg.Name
	}
	return "UNKNOWN"
}

func (l Level) Icon() string {
	if cfg, ok := DefaultLevelConfigs[l]; ok {
		return cfg.Icon
	}
	return "?"
}

func (l Level) Color() color.RGB {
	if cfg, ok := DefaultLevelConfigs[l]; ok {
		return cfg.Color
	}
	return color.RGB{R: 255, G: 255, B: 255}
}

func ParseLevel(s string) Level {
	switch s {
	case "trace", "TRACE":
		return TraceLevel
	case "debug", "DEBUG":
		return DebugLevel
	case "info", "INFO":
		return InfoLevel
	case "success", "SUCCESS":
		return SuccessLevel
	case "warn", "WARN", "warning", "WARNING":
		return WarnLevel
	case "error", "ERROR":
		return ErrorLevel
	case "fatal", "FATAL":
		return FatalLevel
	case "panic", "PANIC":
		return PanicLevel
	default:
		return InfoLevel
	}
}
