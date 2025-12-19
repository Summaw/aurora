package aurora

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/Summaw/aurora/pkg/color"
)

type Logger struct {
	config *Config
	mu     sync.Mutex
	fields []Field
}

func (l *Logger) SetLevel(level Level) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config.Level = level
}

func (l *Logger) SetOutput(w io.Writer) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config.Output = w
}

func (l *Logger) SetTimeFormat(tf string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config.TimeFormat = tf
}

func (l *Logger) EnableCaller(enabled bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config.ShowCaller = enabled
}

func (l *Logger) newEntry(level Level, msg string) *Entry {
	if level < l.config.Level {
		return &Entry{logger: l, discard: true}
	}

	entry := &Entry{
		logger:    l,
		Level:     level,
		Message:   msg,
		Timestamp: time.Now(),
		Fields:    append([]Field{}, l.fields...),
		discard:   false,
	}

	if l.config.ShowCaller {
		entry.Caller = getCaller(l.config.CallerDepth)
	}

	return entry
}

func (l *Logger) Trace(msg string) *Entry {
	return l.newEntry(TraceLevel, msg)
}

func (l *Logger) Debug(msg string) *Entry {
	return l.newEntry(DebugLevel, msg)
}

func (l *Logger) Info(msg string) *Entry {
	return l.newEntry(InfoLevel, msg)
}

func (l *Logger) Success(msg string) *Entry {
	return l.newEntry(SuccessLevel, msg)
}

func (l *Logger) Warn(msg string) *Entry {
	return l.newEntry(WarnLevel, msg)
}

func (l *Logger) Error(msg string) *Entry {
	return l.newEntry(ErrorLevel, msg)
}

func (l *Logger) Fatal(msg string) *Entry {
	entry := l.newEntry(FatalLevel, msg)
	entry.fatal = true
	return entry
}

func (l *Logger) Panic(msg string) *Entry {
	entry := l.newEntry(PanicLevel, msg)
	entry.doPanic = true
	return entry
}

func (l *Logger) WithFields(fields F) *Entry {
	entry := l.newEntry(InfoLevel, "")
	for k, v := range fields {
		entry.Fields = append(entry.Fields, Field{Key: k, Value: v})
	}
	return entry
}

func (l *Logger) With(key string, value any) *Logger {
	newLogger := &Logger{
		config: l.config,
		fields: append(l.fields, Field{Key: key, Value: value}),
	}
	return newLogger
}

func (l *Logger) Ctx(ctx context.Context) *Logger {
	newLogger := &Logger{
		config: l.config,
		fields: append([]Field{}, l.fields...),
	}

	if reqID := ctx.Value("request_id"); reqID != nil {
		newLogger.fields = append(newLogger.fields, Field{Key: "request_id", Value: reqID})
	}
	if traceID := ctx.Value("trace_id"); traceID != nil {
		newLogger.fields = append(newLogger.fields, Field{Key: "trace_id", Value: traceID})
	}

	return newLogger
}

func (l *Logger) write(entry *Entry) {
	if entry.discard {
		return
	}

	l.mu.Lock()
	defer l.mu.Unlock()

	var output string
	if l.config.JSONOutput {
		output = l.formatJSON(entry)
	} else {
		output = l.formatPretty(entry)
	}

	l.config.Output.Write([]byte(output))

	if entry.fatal {
		os.Exit(1)
	}

	if entry.doPanic {
		panic(entry.Message)
	}
}

func (l *Logger) formatPretty(entry *Entry) string {
	var sb strings.Builder

	levelCfg := DefaultLevelConfigs[entry.Level]
	timeStr := entry.Timestamp.Format(l.config.TimeFormat)

	sb.WriteString("\n  ")
	sb.WriteString(color.Colorize(timeStr, color.DimGray))
	sb.WriteString("  ")

	if levelCfg.Bold {
		sb.WriteString(color.ColorizeBold(levelCfg.Icon+" "+levelCfg.Name, levelCfg.Color))
	} else {
		sb.WriteString(color.Colorize(levelCfg.Icon+" "+levelCfg.Name, levelCfg.Color))
	}

	sb.WriteString("  ")
	sb.WriteString(entry.Message)
	sb.WriteString("\n")

	if len(entry.Fields) > 0 {
		for i, field := range entry.Fields {
			prefix := "├─"
			if i == len(entry.Fields)-1 && entry.Caller == "" {
				prefix = "└─"
			}

			sb.WriteString("            ")
			sb.WriteString(color.Colorize(prefix, color.DimGray))
			sb.WriteString(" ")
			sb.WriteString(color.Colorize(field.Key+":", color.Gray))
			sb.WriteString(" ")
			sb.WriteString(fmt.Sprintf("%v", field.Value))
			sb.WriteString("\n")
		}
	}

	if entry.Caller != "" {
		sb.WriteString("            ")
		sb.WriteString(color.Colorize("└─", color.DimGray))
		sb.WriteString(" ")
		sb.WriteString(color.Colorize("at:", color.Gray))
		sb.WriteString(" ")
		sb.WriteString(color.Colorize(entry.Caller, color.DimGray))
		sb.WriteString("\n")
	}

	return sb.String()
}

func (l *Logger) formatJSON(entry *Entry) string {
	var sb strings.Builder

	sb.WriteString(`{"timestamp":"`)
	sb.WriteString(entry.Timestamp.Format(time.RFC3339Nano))
	sb.WriteString(`","level":"`)
	sb.WriteString(entry.Level.String())
	sb.WriteString(`","message":"`)
	sb.WriteString(escapeJSON(entry.Message))
	sb.WriteString(`"`)

	for _, field := range entry.Fields {
		sb.WriteString(`,"`)
		sb.WriteString(field.Key)
		sb.WriteString(`":`)
		sb.WriteString(formatJSONValue(field.Value))
	}

	if entry.Caller != "" {
		sb.WriteString(`,"caller":"`)
		sb.WriteString(entry.Caller)
		sb.WriteString(`"`)
	}

	sb.WriteString("}\n")
	return sb.String()
}

func escapeJSON(s string) string {
	s = strings.ReplaceAll(s, `\`, `\\`)
	s = strings.ReplaceAll(s, `"`, `\"`)
	s = strings.ReplaceAll(s, "\n", `\n`)
	s = strings.ReplaceAll(s, "\r", `\r`)
	s = strings.ReplaceAll(s, "\t", `\t`)
	return s
}

func formatJSONValue(v any) string {
	switch val := v.(type) {
	case string:
		return `"` + escapeJSON(val) + `"`
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val)
	case float32, float64:
		return fmt.Sprintf("%g", val)
	case bool:
		return fmt.Sprintf("%t", val)
	default:
		return `"` + escapeJSON(fmt.Sprintf("%v", val)) + `"`
	}
}

func getCaller(depth int) string {
	_, file, line, ok := runtime.Caller(depth)
	if !ok {
		return ""
	}

	parts := strings.Split(file, "/")
	if len(parts) > 0 {
		file = parts[len(parts)-1]
	}

	return fmt.Sprintf("%s:%d", file, line)
}
