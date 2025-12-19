package color

import (
	"fmt"
	"strconv"
	"strings"
)

type RGB struct {
	R, G, B uint8
}

func (c RGB) ANSI() string {
	return fmt.Sprintf("\x1b[38;2;%d;%d;%dm", c.R, c.G, c.B)
}

func (c RGB) ANSIBg() string {
	return fmt.Sprintf("\x1b[48;2;%d;%d;%dm", c.R, c.G, c.B)
}

func (c RGB) Hex() string {
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

func (c RGB) Blend(other RGB, t float64) RGB {
	return RGB{
		R: uint8(float64(c.R) + t*(float64(other.R)-float64(c.R))),
		G: uint8(float64(c.G) + t*(float64(other.G)-float64(c.G))),
		B: uint8(float64(c.B) + t*(float64(other.B)-float64(c.B))),
	}
}

func FromHex(hex string) RGB {
	hex = strings.TrimPrefix(hex, "#")
	if len(hex) == 3 {
		hex = string(hex[0]) + string(hex[0]) + string(hex[1]) + string(hex[1]) + string(hex[2]) + string(hex[2])
	}
	if len(hex) != 6 {
		return RGB{255, 255, 255}
	}

	r, _ := strconv.ParseUint(hex[0:2], 16, 8)
	g, _ := strconv.ParseUint(hex[2:4], 16, 8)
	b, _ := strconv.ParseUint(hex[4:6], 16, 8)

	return RGB{uint8(r), uint8(g), uint8(b)}
}

const Reset = "\x1b[0m"
const Bold = "\x1b[1m"
const Dim = "\x1b[2m"
const Italic = "\x1b[3m"
const Underline = "\x1b[4m"

func Colorize(text string, c RGB) string {
	return c.ANSI() + text + Reset
}

func ColorizeBold(text string, c RGB) string {
	return Bold + c.ANSI() + text + Reset
}

func ColorizeWithBg(text string, fg, bg RGB) string {
	return fg.ANSI() + bg.ANSIBg() + text + Reset
}

func ApplyStyle(text string, styles ...string) string {
	prefix := strings.Join(styles, "")
	return prefix + text + Reset
}

var (
	White   = RGB{255, 255, 255}
	Black   = RGB{0, 0, 0}
	Red     = RGB{239, 68, 68}
	Green   = RGB{34, 197, 94}
	Blue    = RGB{59, 130, 246}
	Yellow  = RGB{234, 179, 8}
	Cyan    = RGB{6, 182, 212}
	Magenta = RGB{168, 85, 247}
	Orange  = RGB{249, 115, 22}
	Pink    = RGB{236, 72, 153}
	Gray    = RGB{107, 114, 128}
	DimGray = RGB{75, 85, 99}
)
