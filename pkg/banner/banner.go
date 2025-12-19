package banner

import (
	"fmt"
	"os"
	"strings"

	"github.com/Summaw/aurora/pkg/color"
)

type Builder struct {
	text     string
	font     string
	gradient color.Gradient
	tagline  string
	version  string
	border   string
	padding  int
	width    int
	output   *os.File
}

func New(text string) *Builder {
	return &Builder{
		text:     text,
		font:     "block",
		gradient: color.GradientAurora,
		border:   "rounded",
		padding:  1,
		output:   os.Stdout,
	}
}

func (b *Builder) Font(font string) *Builder {
	b.font = font
	return b
}

func (b *Builder) Gradient(name string) *Builder {
	b.gradient = color.GetGradient(name)
	return b
}

func (b *Builder) GradientRGB(start, end color.RGB) *Builder {
	b.gradient = color.NewGradient(start, end)
	return b
}

func (b *Builder) GradientMulti(colors ...string) *Builder {
	b.gradient = color.NewMultiGradient(colors...)
	return b
}

func (b *Builder) GradientCustom(g color.Gradient) *Builder {
	b.gradient = g
	return b
}

func (b *Builder) Tagline(tagline string) *Builder {
	b.tagline = tagline
	return b
}

func (b *Builder) Version(version string) *Builder {
	b.version = version
	return b
}

func (b *Builder) Border(style string) *Builder {
	b.border = style
	return b
}

func (b *Builder) Padding(p int) *Builder {
	b.padding = p
	return b
}

func (b *Builder) Width(w int) *Builder {
	b.width = w
	return b
}

func (b *Builder) Output(f *os.File) *Builder {
	b.output = f
	return b
}

func (b *Builder) Build() string {
	artLines := GenerateArt(b.text, b.font)
	coloredLines := b.gradient.ApplyLines(artLines)

	maxWidth := 0
	for _, line := range artLines {
		lineLen := len([]rune(line))
		if lineLen > maxWidth {
			maxWidth = lineLen
		}
	}

	if b.tagline != "" && len(b.tagline) > maxWidth {
		maxWidth = len(b.tagline)
	}
	if b.version != "" && len(b.version) > maxWidth {
		maxWidth = len(b.version)
	}

	if b.width > maxWidth {
		maxWidth = b.width
	}

	maxWidth += b.padding * 4

	var result strings.Builder
	borderChars := getBorderChars(b.border)

	result.WriteString("\n")
	result.WriteString(borderChars.topLeft)
	result.WriteString(strings.Repeat(borderChars.horizontal, maxWidth+2))
	result.WriteString(borderChars.topRight)
	result.WriteString("\n")

	for i := 0; i < b.padding; i++ {
		result.WriteString(borderChars.vertical)
		result.WriteString(strings.Repeat(" ", maxWidth+2))
		result.WriteString(borderChars.vertical)
		result.WriteString("\n")
	}

	for idx, line := range coloredLines {
		result.WriteString(borderChars.vertical)
		result.WriteString(" ")
		actualLen := len([]rune(artLines[idx]))
		paddedLine := centerText(line, maxWidth, actualLen)
		result.WriteString(paddedLine)
		result.WriteString(" ")
		result.WriteString(borderChars.vertical)
		result.WriteString("\n")
	}

	for i := 0; i < b.padding; i++ {
		result.WriteString(borderChars.vertical)
		result.WriteString(strings.Repeat(" ", maxWidth+2))
		result.WriteString(borderChars.vertical)
		result.WriteString("\n")
	}

	if b.tagline != "" {
		taglineColored := b.gradient.Apply(b.tagline)
		result.WriteString(borderChars.vertical)
		result.WriteString(" ")
		paddedTagline := centerText(taglineColored, maxWidth, len(b.tagline))
		result.WriteString(paddedTagline)
		result.WriteString(" ")
		result.WriteString(borderChars.vertical)
		result.WriteString("\n")
	}

	if b.version != "" {
		versionColored := color.Colorize(b.version, color.Gray)
		result.WriteString(borderChars.vertical)
		result.WriteString(" ")
		paddedVersion := centerText(versionColored, maxWidth, len(b.version))
		result.WriteString(paddedVersion)
		result.WriteString(" ")
		result.WriteString(borderChars.vertical)
		result.WriteString("\n")
	}

	if b.tagline != "" || b.version != "" {
		result.WriteString(borderChars.vertical)
		result.WriteString(strings.Repeat(" ", maxWidth+2))
		result.WriteString(borderChars.vertical)
		result.WriteString("\n")
	}

	result.WriteString(borderChars.bottomLeft)
	result.WriteString(strings.Repeat(borderChars.horizontal, maxWidth+2))
	result.WriteString(borderChars.bottomRight)
	result.WriteString("\n")

	return result.String()
}

func (b *Builder) Render() {
	fmt.Fprint(b.output, b.Build())
}

func (b *Builder) String() string {
	return b.Build()
}

func centerText(text string, totalWidth int, actualLen int) string {
	padding := (totalWidth - actualLen) / 2
	if padding < 0 {
		padding = 0
	}
	rightPad := totalWidth - actualLen - padding
	if rightPad < 0 {
		rightPad = 0
	}
	return strings.Repeat(" ", padding) + text + strings.Repeat(" ", rightPad)
}

type borderSet struct {
	topLeft     string
	topRight    string
	bottomLeft  string
	bottomRight string
	horizontal  string
	vertical    string
}

func getBorderChars(style string) borderSet {
	switch style {
	case "rounded":
		return borderSet{"╭", "╮", "╰", "╯", "─", "│"}
	case "sharp":
		return borderSet{"┌", "┐", "└", "┘", "─", "│"}
	case "double":
		return borderSet{"╔", "╗", "╚", "╝", "═", "║"}
	case "heavy":
		return borderSet{"┏", "┓", "┗", "┛", "━", "┃"}
	case "ascii":
		return borderSet{"+", "+", "+", "+", "-", "|"}
	case "none":
		return borderSet{" ", " ", " ", " ", " ", " "}
	default:
		return borderSet{"╭", "╮", "╰", "╯", "─", "│"}
	}
}
