package aurora

import (
	"fmt"
	"time"
)

type Field struct {
	Key   string
	Value any
}

type F map[string]any

type Entry struct {
	logger    *Logger
	Level     Level
	Message   string
	Timestamp time.Time
	Fields    []Field
	Caller    string
	discard   bool
	fatal     bool
	doPanic   bool
}

func (e *Entry) Str(key, value string) *Entry {
	if e.discard {
		return e
	}
	e.Fields = append(e.Fields, Field{Key: key, Value: value})
	return e
}

func (e *Entry) Int(key string, value int) *Entry {
	if e.discard {
		return e
	}
	e.Fields = append(e.Fields, Field{Key: key, Value: value})
	return e
}

func (e *Entry) Int64(key string, value int64) *Entry {
	if e.discard {
		return e
	}
	e.Fields = append(e.Fields, Field{Key: key, Value: value})
	return e
}

func (e *Entry) Uint(key string, value uint) *Entry {
	if e.discard {
		return e
	}
	e.Fields = append(e.Fields, Field{Key: key, Value: value})
	return e
}

func (e *Entry) Uint64(key string, value uint64) *Entry {
	if e.discard {
		return e
	}
	e.Fields = append(e.Fields, Field{Key: key, Value: value})
	return e
}

func (e *Entry) Float32(key string, value float32) *Entry {
	if e.discard {
		return e
	}
	e.Fields = append(e.Fields, Field{Key: key, Value: value})
	return e
}

func (e *Entry) Float64(key string, value float64) *Entry {
	if e.discard {
		return e
	}
	e.Fields = append(e.Fields, Field{Key: key, Value: value})
	return e
}

func (e *Entry) Bool(key string, value bool) *Entry {
	if e.discard {
		return e
	}
	e.Fields = append(e.Fields, Field{Key: key, Value: value})
	return e
}

func (e *Entry) Dur(key string, value time.Duration) *Entry {
	if e.discard {
		return e
	}
	e.Fields = append(e.Fields, Field{Key: key, Value: formatDuration(value)})
	return e
}

func (e *Entry) Time(key string, value time.Time) *Entry {
	if e.discard {
		return e
	}
	e.Fields = append(e.Fields, Field{Key: key, Value: value.Format(time.RFC3339)})
	return e
}

func (e *Entry) Any(key string, value any) *Entry {
	if e.discard {
		return e
	}
	e.Fields = append(e.Fields, Field{Key: key, Value: value})
	return e
}

func (e *Entry) Err(err error) *Entry {
	if e.discard || err == nil {
		return e
	}
	e.Fields = append(e.Fields, Field{Key: "error", Value: err.Error()})
	return e
}

func (e *Entry) Field(key string, value any) *Entry {
	if e.discard {
		return e
	}
	e.Fields = append(e.Fields, Field{Key: key, Value: value})
	return e
}

func (e *Entry) WithFields(fields F) *Entry {
	if e.discard {
		return e
	}
	for k, v := range fields {
		e.Fields = append(e.Fields, Field{Key: k, Value: v})
	}
	return e
}

func (e *Entry) Msg(msg string) {
	if e.discard {
		return
	}
	e.Message = msg
	e.logger.write(e)
}

func (e *Entry) Msgf(format string, args ...any) {
	if e.discard {
		return
	}
	e.Message = fmt.Sprintf(format, args...)
	e.logger.write(e)
}

func (e *Entry) Send() {
	if e.discard {
		return
	}
	e.logger.write(e)
}

func formatDuration(d time.Duration) string {
	if d < time.Microsecond {
		return fmt.Sprintf("%dns", d.Nanoseconds())
	}
	if d < time.Millisecond {
		return fmt.Sprintf("%.2fµs", float64(d.Nanoseconds())/1000)
	}
	if d < time.Second {
		return fmt.Sprintf("%.2fms", float64(d.Nanoseconds())/1e6)
	}
	if d < time.Minute {
		return fmt.Sprintf("%.2fs", d.Seconds())
	}
	return d.String()
}
