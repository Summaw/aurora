package aurora

import (
	"io"
	"os"
	"sync"

	"github.com/Summaw/aurora/pkg/banner"
	"github.com/Summaw/aurora/pkg/color"
	"github.com/Summaw/aurora/pkg/style"
)

var (
	std     *Logger
	stdOnce sync.Once
)

func Default() *Logger {
	stdOnce.Do(func() {
		std = New()
	})
	return std
}

func New(opts ...Option) *Logger {
	cfg := &Config{
		Output:      os.Stdout,
		Level:       InfoLevel,
		TimeFormat:  "15:04:05.000",
		ShowCaller:  false,
		CallerDepth: 4,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	return &Logger{
		config: cfg,
		mu:     sync.Mutex{},
	}
}

func Banner(text string) *banner.Builder {
	return banner.New(text)
}

func Gradient(name string) color.Gradient {
	return color.GetGradient(name)
}

func GradientRGB(start, end color.RGB) color.Gradient {
	return color.NewGradient(start, end)
}

func GradientMulti(colors ...string) color.Gradient {
	return color.NewMultiGradient(colors...)
}

func RGB(r, g, b uint8) color.RGB {
	return color.RGB{R: r, G: g, B: b}
}

func Hex(hex string) color.RGB {
	return color.FromHex(hex)
}

func Box(content string) *style.BoxBuilder {
	return style.NewBox(content)
}

func Table(headers []string, rows [][]string) *style.TableBuilder {
	return style.NewTable(headers, rows)
}

func Divider(text string) *style.DividerBuilder {
	return style.NewDivider(text)
}

func KV(pairs map[string]any) *style.KVBuilder {
	return style.NewKV(pairs)
}

func Spin(message string) *style.Spinner {
	return style.NewSpinner(message)
}

func Progress(label string, total int) *style.ProgressBar {
	return style.NewProgressBar(label, total)
}

func Trace(msg string) *Entry {
	return Default().Trace(msg)
}

func Debug(msg string) *Entry {
	return Default().Debug(msg)
}

func Info(msg string) *Entry {
	return Default().Info(msg)
}

func Success(msg string) *Entry {
	return Default().Success(msg)
}

func Warn(msg string) *Entry {
	return Default().Warn(msg)
}

func Error(msg string) *Entry {
	return Default().Error(msg)
}

func Fatal(msg string) *Entry {
	return Default().Fatal(msg)
}

func SetLevel(level Level) {
	Default().SetLevel(level)
}

func SetOutput(w io.Writer) {
	Default().SetOutput(w)
}

func WithFields(fields F) *Entry {
	return Default().WithFields(fields)
}
