package aurora

import "io"

type Option func(*Config)

func WithOutput(w io.Writer) Option {
	return func(c *Config) {
		c.Output = w
	}
}

func WithLevel(level Level) Option {
	return func(c *Config) {
		c.Level = level
	}
}

func WithTimeFormat(tf string) Option {
	return func(c *Config) {
		c.TimeFormat = tf
	}
}

func WithCaller(enabled bool) Option {
	return func(c *Config) {
		c.ShowCaller = enabled
	}
}

func WithCallerDepth(depth int) Option {
	return func(c *Config) {
		c.CallerDepth = depth
	}
}

func WithJSON(enabled bool) Option {
	return func(c *Config) {
		c.JSONOutput = enabled
	}
}
