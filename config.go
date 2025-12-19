package aurora

import "io"

type Config struct {
	Output      io.Writer
	Level       Level
	TimeFormat  string
	ShowCaller  bool
	CallerDepth int
	JSONOutput  bool
}
