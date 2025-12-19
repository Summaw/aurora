# Aurora API Reference

## Package aurora

### Logger Creation

#### New
```go
func New(opts ...Option) *Logger
```
Creates a new Logger instance.

**Options:**
- `WithLevel(level Level)` - Set minimum log level
- `WithOutput(w io.Writer)` - Set output destination
- `WithCaller(enabled bool)` - Enable caller info (file:line)
- `WithTimeFormat(format string)` - Set time format
- `WithJSON(enabled bool)` - Enable JSON output

**Example:**
```go
log := aurora.New(
    aurora.WithLevel(aurora.DebugLevel),
    aurora.WithCaller(true),
)
```

#### Default
```go
func Default() *Logger
```
Returns the global default logger.

---

### Log Levels

```go
const (
    TraceLevel   Level = iota  // Most verbose
    DebugLevel                  // Debug info
    InfoLevel                   // General info
    SuccessLevel               // Success messages
    WarnLevel                  // Warnings
    ErrorLevel                 // Errors
    FatalLevel                 // Fatal (exits)
    PanicLevel                 // Panic
    Disabled                   // Disable logging
)
```

---

### Logging Methods

```go
func (l *Logger) Trace(msg string) *Entry
func (l *Logger) Debug(msg string) *Entry
func (l *Logger) Info(msg string) *Entry
func (l *Logger) Success(msg string) *Entry
func (l *Logger) Warn(msg string) *Entry
func (l *Logger) Error(msg string) *Entry
func (l *Logger) Fatal(msg string) *Entry  // Calls os.Exit(1)
func (l *Logger) Panic(msg string) *Entry  // Calls panic()
```

---

### Entry Field Methods

All return `*Entry` for chaining.

```go
func (e *Entry) Str(key, value string) *Entry
func (e *Entry) Int(key string, value int) *Entry
func (e *Entry) Int64(key string, value int64) *Entry
func (e *Entry) Uint(key string, value uint) *Entry
func (e *Entry) Float64(key string, value float64) *Entry
func (e *Entry) Bool(key string, value bool) *Entry
func (e *Entry) Dur(key string, value time.Duration) *Entry
func (e *Entry) Time(key string, value time.Time) *Entry
func (e *Entry) Any(key string, value any) *Entry
func (e *Entry) Err(err error) *Entry
func (e *Entry) WithFields(fields F) *Entry
```

#### Finalizers
```go
func (e *Entry) Send()                           // Write entry
func (e *Entry) Msg(msg string)                  // Set message and write
func (e *Entry) Msgf(format string, args ...any) // Formatted message
```

---

### Banner

```go
func Banner(text string) *banner.Builder
```

**Builder Methods:**
```go
func (b *Builder) Font(font string) *Builder         // block, slant, minimal
func (b *Builder) Gradient(name string) *Builder     // Preset gradient name
func (b *Builder) GradientRGB(start, end RGB) *Builder
func (b *Builder) GradientMulti(colors ...string) *Builder
func (b *Builder) Tagline(text string) *Builder
func (b *Builder) Version(text string) *Builder
func (b *Builder) Border(style string) *Builder      // rounded, sharp, double, heavy, ascii
func (b *Builder) Padding(p int) *Builder
func (b *Builder) Render()                           // Output to stdout
func (b *Builder) Build() string                     // Return as string
```

---

### UI Components

#### Table
```go
func Table(headers []string, rows [][]string) *style.TableBuilder

func (t *TableBuilder) Border(style string) *TableBuilder
func (t *TableBuilder) Gradient(name string) *TableBuilder
func (t *TableBuilder) Render()
```

#### Divider
```go
func Divider(text string) *style.DividerBuilder

func (d *DividerBuilder) Width(w int) *DividerBuilder
func (d *DividerBuilder) Char(c string) *DividerBuilder
func (d *DividerBuilder) Gradient(name string) *DividerBuilder
func (d *DividerBuilder) Render()
```

#### KV (Key-Value)
```go
func KV(pairs map[string]any) *style.KVBuilder

func (k *KVBuilder) Gradient(name string) *KVBuilder
func (k *KVBuilder) Render()
```

#### Box
```go
func Box(content string) *style.BoxBuilder

func (b *BoxBuilder) Title(t string) *BoxBuilder
func (b *BoxBuilder) Border(style string) *BoxBuilder
func (b *BoxBuilder) Padding(p int) *BoxBuilder
func (b *BoxBuilder) Gradient(name string) *BoxBuilder
func (b *BoxBuilder) Render()
```

#### Spinner
```go
func Spin(message string) *style.Spinner

func (s *Spinner) Success(msg string)  // Stop with ✓
func (s *Spinner) Fail(msg string)     // Stop with ✖
func (s *Spinner) Warn(msg string)     // Stop with ⚠
func (s *Spinner) Info(msg string)     // Stop with ●
func (s *Spinner) Stop()               // Stop silently
```

#### Progress Bar
```go
func Progress(label string, total int) *style.ProgressBar

func (p *ProgressBar) Width(w int) *ProgressBar
func (p *ProgressBar) Gradient(name string) *ProgressBar
func (p *ProgressBar) Chars(complete, pending string) *ProgressBar
func (p *ProgressBar) Set(value int)
func (p *ProgressBar) Increment()
func (p *ProgressBar) Done()
```

---

### Color Utilities

```go
func RGB(r, g, b uint8) color.RGB
func Hex(hex string) color.RGB
func Gradient(name string) color.Gradient
func GradientRGB(start, end color.RGB) color.Gradient
func GradientMulti(colors ...string) color.Gradient
```

---

### Available Gradients

`aurora`, `sunset`, `ocean`, `neon`, `cyberpunk`, `miami`, `fire`, `forest`, `galaxy`, `retro`, `mint`, `peach`, `lavender`, `gold`, `ice`, `blood`, `matrix`, `vaporwave`, `rainbow`, `terminal`, `rose`, `sky`

---

### Types

```go
type F map[string]any  // Field map shorthand

type Level int         // Log level

type Field struct {
    Key   string
    Value any
}
```
